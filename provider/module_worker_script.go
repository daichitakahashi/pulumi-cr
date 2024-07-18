package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	ModuleWorkerScript struct{}

	ModuleWorkerScriptArgs struct {
		AccountID               string                   `pulumi:"accountId"`
		Name                    string                   `pulumi:"name"`
		Dir                     string                   `pulumi:"dir"`
		MainModule              string                   `pulumi:"mainModule"`
		CompatibilityDate       string                   `pulumi:"compatibilityDate,optional"`
		CompatibilityFlags      []string                 `pulumi:"compatibilityFlags,optional"`
		PlainTextBindings       []PlainTextBinding       `pulumi:"plainTextBindings,optional"`
		SecretTextBindings      []SecretTextBinding      `pulumi:"secretTextBindings,optional"`
		KVNamespaceBindings     []KVNamespaceBinding     `pulumi:"kvNamespaceBindings,optional"`
		R2BucketBindings        []R2BucketBinding        `pulumi:"r2BucketBindings,optional"`
		QueueBindings           []QueueBinding           `pulumi:"queueBindings,optional"`
		D1DatabaseBindings      []D1DatabaseBinding      `pulumi:"d1DatabaseBindings,optional"`
		ServiceBindings         []ServiceBinding         `pulumi:"serviceBindings,optional"`
		AnalyticsEngineBindings []AnalyticsEngineBinding `pulumi:"analyticsEngineBindings,optional"`
		Logpush                 bool                     `pulumi:"logpush,optional"`
	}

	PlainTextBinding struct {
		Name string `pulumi:"name"`
		Text string `pulumi:"text"`
	}
	SecretTextBinding struct {
		Name string `pulumi:"name"`
		Text string `pulumi:"text"`
	}
	KVNamespaceBinding struct {
		Name        string `pulumi:"name"`
		NamespaceID string `pulumi:"namespaceId"`
	}
	R2BucketBinding struct {
		Name       string `pulumi:"name"`
		BucketName string `pulumi:"bucketName"`
	}
	QueueBinding struct {
		Name      string `pulumi:"name"`
		QueueName string `pulumi:"queueName"`
	}
	D1DatabaseBinding struct {
		Name       string `pulumi:"name"`
		DatabaseID string `pulumi:"databaseId"`
	}
	ServiceBinding struct {
		Name    string `pulumi:"name"`
		Service string `pulumi:"service"`
	}
	AnalyticsEngineBinding struct {
		Name    string `pulumi:"name"`
		Dataset string `pulumi:"dataset"`
	}

	ModuleWorkerScriptState struct {
		ModuleWorkerScriptArgs
		ETag string `pulumi:"eTag"`
	}
)

// Create implements infer.CustomResource.
func (m *ModuleWorkerScript) Create(ctx p.Context, name string, args ModuleWorkerScriptArgs, preview bool) (id string, output ModuleWorkerScriptState, err error) {
	output.ModuleWorkerScriptArgs = args
	if preview {
		return name, output, nil
	}
	eTag, err := uploadModuleWorkerScript(ctx, args)
	if err != nil {
		return name, output, err
	}
	output.ETag = eTag
	return name, output, nil
}

// Update implements infer.CustomUpdate.
func (m *ModuleWorkerScript) Update(ctx p.Context, id string, olds ModuleWorkerScriptState, news ModuleWorkerScriptArgs, preview bool) (ModuleWorkerScriptState, error) {
	notChanged := reflect.DeepEqual(olds.ModuleWorkerScriptArgs, news)
	if notChanged {
		return olds, nil
	}
	output := ModuleWorkerScriptState{
		ModuleWorkerScriptArgs: news,
		ETag:                   "",
	}
	if preview {
		return output, nil
	}

	eTag, err := uploadModuleWorkerScript(ctx, news)
	if err != nil {
		return output, err
	}
	output.ETag = eTag
	return output, nil
}

// Delete implements infer.CustomDelete.
func (m *ModuleWorkerScript) Delete(ctx p.Context, id string, props ModuleWorkerScriptState) error {
	cli := getConfig(ctx).client
	return cli.Workers.Scripts.Delete(ctx, props.Name, workers.ScriptDeleteParams{
		AccountID: cloudflare.String(props.AccountID),
	})
}

var _ infer.CustomResource[ModuleWorkerScriptArgs, ModuleWorkerScriptState] = (*ModuleWorkerScript)(nil)
var _ infer.CustomUpdate[ModuleWorkerScriptArgs, ModuleWorkerScriptState] = (*ModuleWorkerScript)(nil)
var _ infer.CustomDelete[ModuleWorkerScriptState] = (*ModuleWorkerScript)(nil)

func appendPart(w *multipart.Writer, filename, contentType string) error {
	name := filepath.Base(filename)
	part, err := w.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": {fmt.Sprintf("form-data; name=%q; filename=%q", name, name)},
		"Content-Type":        {contentType},
	})
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(part, file)
	return err
}

var (
	javascriptRx = regexp.MustCompile(`(?i)^.*(\.js|\.mjs)$`)
	wasmRx       = regexp.MustCompile(`(?i)^.*\.wasm$`)
	sourceMapRx  = regexp.MustCompile(`(?i)^.*\.js\.map$`)
)

func uploadModuleWorkerScript(ctx p.Context, input ModuleWorkerScriptArgs) (etag string, _ error) {
	config := getConfig(ctx)
	if input.AccountID == "" {
		return "", errors.New("empty accountId provided")
	}
	if input.Name == "" {
		return "", errors.New("empty name provided")
	}
	if input.Dir == "" {
		return "", errors.New("empty dir provided")
	}

	// List and extract applicable files in `input.Dir`.
	var (
		fsys        = os.DirFS(input.Dir)
		scriptFiles [][2]string // {filename, contentType}
	)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		switch {
		case javascriptRx.MatchString(path):
			scriptFiles = append(scriptFiles, [2]string{path, "application/javascript+module"})
		case wasmRx.MatchString(path):
			scriptFiles = append(scriptFiles, [2]string{path, "application/wasm"})
		case sourceMapRx.MatchString(path):
			scriptFiles = append(scriptFiles, [2]string{path, "application/source-map"})
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	// Check if `input.MailModule` is present.
	found := contains(scriptFiles, func(f [2]string) bool {
		return f[0] == input.MainModule
	})
	if !found {
		return "", fmt.Errorf("mail module %q not found", input.MainModule)
	}

	buf := bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)

	// Write meta part.
	metaPart, err := w.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": {`form-data; name="metadata"`},
		"Content-Type":        {"application/json"},
	})
	if err != nil {
		return "", err
	}

	var bindings []map[string]string
	for _, plainText := range input.PlainTextBindings {
		bindings = append(bindings, map[string]string{
			"type": "plain_text",
			"name": plainText.Name,
			"text": plainText.Text,
		})
	}
	for _, secretText := range input.SecretTextBindings {
		bindings = append(bindings, map[string]string{
			"type": "secret_text",
			"name": secretText.Name,
			"text": secretText.Text,
		})
	}
	for _, kvNamespace := range input.KVNamespaceBindings {
		bindings = append(bindings, map[string]string{
			"type":         "kv_namespace",
			"name":         kvNamespace.Name,
			"namespace_id": kvNamespace.NamespaceID,
		})
	}
	for _, r2Bucket := range input.R2BucketBindings {
		bindings = append(bindings, map[string]string{
			"type":        "r2_bucket",
			"name":        r2Bucket.Name,
			"bucket_name": r2Bucket.BucketName,
		})
	}
	for _, queue := range input.QueueBindings {
		bindings = append(bindings, map[string]string{
			"type":       "queue",
			"name":       queue.Name,
			"queue_name": queue.QueueName,
		})
	}
	for _, d1Database := range input.D1DatabaseBindings {
		bindings = append(bindings, map[string]string{
			"type": "d1",
			"name": d1Database.Name,
			"id":   d1Database.DatabaseID,
		})
	}
	for _, service := range input.ServiceBindings {
		bindings = append(bindings, map[string]string{
			"type":    "service",
			"name":    service.Name,
			"service": service.Service,
		})
	}
	for _, analyticsEngine := range input.AnalyticsEngineBindings {
		bindings = append(bindings, map[string]string{
			"type":    "analytics_engine",
			"name":    analyticsEngine.Name,
			"dataset": analyticsEngine.Dataset,
		})
	}
	metadata := map[string]any{
		"main_module": input.MainModule,
		"bindings":    bindings,
		"logpush":     input.Logpush,
	}
	if input.CompatibilityDate != "" {
		metadata["compatibility_date"] = input.CompatibilityDate
	}
	if len(input.CompatibilityFlags) > 0 {
		metadata["compatibility_flags"] = input.CompatibilityFlags
	}

	err = json.NewEncoder(metaPart).Encode(metadata)
	if err != nil {
		return "", err
	}

	// Write script files.
	for _, e := range scriptFiles {
		err := appendPart(w, e[0], e[1])
		if err != nil {
			return "", err
		}
	}
	w.Close()

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/workers/scripts/%s", url.PathEscape(input.AccountID), url.PathEscape(input.Name)),
		buf,
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+config.CloudflareAPIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()
	var e workers.ScriptUpdateResponseEnvelope
	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		return "", err
	}

	if !e.Success {
		return "", errors.New(e.Errors[0].Message)
	}

	return e.Result.Etag, nil
}
