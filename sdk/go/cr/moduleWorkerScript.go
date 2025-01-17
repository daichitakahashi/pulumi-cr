// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"internal"
)

type ModuleWorkerScript struct {
	pulumi.CustomResourceState

	AccountId               pulumi.StringOutput               `pulumi:"accountId"`
	AnalyticsEngineBindings AnalyticsEngineBindingArrayOutput `pulumi:"analyticsEngineBindings"`
	CompatibilityDate       pulumi.StringPtrOutput            `pulumi:"compatibilityDate"`
	CompatibilityFlags      pulumi.StringArrayOutput          `pulumi:"compatibilityFlags"`
	D1DatabaseBindings      D1DatabaseBindingArrayOutput      `pulumi:"d1DatabaseBindings"`
	Dir                     pulumi.StringOutput               `pulumi:"dir"`
	ETag                    pulumi.StringOutput               `pulumi:"eTag"`
	KvNamespaceBindings     KVNamespaceBindingArrayOutput     `pulumi:"kvNamespaceBindings"`
	Logpush                 pulumi.BoolPtrOutput              `pulumi:"logpush"`
	MainModule              pulumi.StringOutput               `pulumi:"mainModule"`
	Name                    pulumi.StringOutput               `pulumi:"name"`
	PlainTextBindings       PlainTextBindingArrayOutput       `pulumi:"plainTextBindings"`
	QueueBindings           QueueBindingArrayOutput           `pulumi:"queueBindings"`
	R2BucketBindings        R2BucketBindingArrayOutput        `pulumi:"r2BucketBindings"`
	SecretTextBindings      SecretTextBindingArrayOutput      `pulumi:"secretTextBindings"`
	ServiceBindings         ServiceBindingArrayOutput         `pulumi:"serviceBindings"`
}

// NewModuleWorkerScript registers a new resource with the given unique name, arguments, and options.
func NewModuleWorkerScript(ctx *pulumi.Context,
	name string, args *ModuleWorkerScriptArgs, opts ...pulumi.ResourceOption) (*ModuleWorkerScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.Dir == nil {
		return nil, errors.New("invalid value for required argument 'Dir'")
	}
	if args.MainModule == nil {
		return nil, errors.New("invalid value for required argument 'MainModule'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ModuleWorkerScript
	err := ctx.RegisterResource("cr:index:ModuleWorkerScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModuleWorkerScript gets an existing ModuleWorkerScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModuleWorkerScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleWorkerScriptState, opts ...pulumi.ResourceOption) (*ModuleWorkerScript, error) {
	var resource ModuleWorkerScript
	err := ctx.ReadResource("cr:index:ModuleWorkerScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModuleWorkerScript resources.
type moduleWorkerScriptState struct {
}

type ModuleWorkerScriptState struct {
}

func (ModuleWorkerScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleWorkerScriptState)(nil)).Elem()
}

type moduleWorkerScriptArgs struct {
	AccountId               string                   `pulumi:"accountId"`
	AnalyticsEngineBindings []AnalyticsEngineBinding `pulumi:"analyticsEngineBindings"`
	CompatibilityDate       *string                  `pulumi:"compatibilityDate"`
	CompatibilityFlags      []string                 `pulumi:"compatibilityFlags"`
	D1DatabaseBindings      []D1DatabaseBinding      `pulumi:"d1DatabaseBindings"`
	Dir                     string                   `pulumi:"dir"`
	KvNamespaceBindings     []KVNamespaceBinding     `pulumi:"kvNamespaceBindings"`
	Logpush                 *bool                    `pulumi:"logpush"`
	MainModule              string                   `pulumi:"mainModule"`
	Name                    string                   `pulumi:"name"`
	PlainTextBindings       []PlainTextBinding       `pulumi:"plainTextBindings"`
	QueueBindings           []QueueBinding           `pulumi:"queueBindings"`
	R2BucketBindings        []R2BucketBinding        `pulumi:"r2BucketBindings"`
	SecretTextBindings      []SecretTextBinding      `pulumi:"secretTextBindings"`
	ServiceBindings         []ServiceBinding         `pulumi:"serviceBindings"`
}

// The set of arguments for constructing a ModuleWorkerScript resource.
type ModuleWorkerScriptArgs struct {
	AccountId               pulumi.StringInput
	AnalyticsEngineBindings AnalyticsEngineBindingArrayInput
	CompatibilityDate       pulumi.StringPtrInput
	CompatibilityFlags      pulumi.StringArrayInput
	D1DatabaseBindings      D1DatabaseBindingArrayInput
	Dir                     pulumi.StringInput
	KvNamespaceBindings     KVNamespaceBindingArrayInput
	Logpush                 pulumi.BoolPtrInput
	MainModule              pulumi.StringInput
	Name                    pulumi.StringInput
	PlainTextBindings       PlainTextBindingArrayInput
	QueueBindings           QueueBindingArrayInput
	R2BucketBindings        R2BucketBindingArrayInput
	SecretTextBindings      SecretTextBindingArrayInput
	ServiceBindings         ServiceBindingArrayInput
}

func (ModuleWorkerScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleWorkerScriptArgs)(nil)).Elem()
}

type ModuleWorkerScriptInput interface {
	pulumi.Input

	ToModuleWorkerScriptOutput() ModuleWorkerScriptOutput
	ToModuleWorkerScriptOutputWithContext(ctx context.Context) ModuleWorkerScriptOutput
}

func (*ModuleWorkerScript) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleWorkerScript)(nil)).Elem()
}

func (i *ModuleWorkerScript) ToModuleWorkerScriptOutput() ModuleWorkerScriptOutput {
	return i.ToModuleWorkerScriptOutputWithContext(context.Background())
}

func (i *ModuleWorkerScript) ToModuleWorkerScriptOutputWithContext(ctx context.Context) ModuleWorkerScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleWorkerScriptOutput)
}

type ModuleWorkerScriptOutput struct{ *pulumi.OutputState }

func (ModuleWorkerScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleWorkerScript)(nil)).Elem()
}

func (o ModuleWorkerScriptOutput) ToModuleWorkerScriptOutput() ModuleWorkerScriptOutput {
	return o
}

func (o ModuleWorkerScriptOutput) ToModuleWorkerScriptOutputWithContext(ctx context.Context) ModuleWorkerScriptOutput {
	return o
}

func (o ModuleWorkerScriptOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o ModuleWorkerScriptOutput) AnalyticsEngineBindings() AnalyticsEngineBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) AnalyticsEngineBindingArrayOutput { return v.AnalyticsEngineBindings }).(AnalyticsEngineBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) CompatibilityDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringPtrOutput { return v.CompatibilityDate }).(pulumi.StringPtrOutput)
}

func (o ModuleWorkerScriptOutput) CompatibilityFlags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringArrayOutput { return v.CompatibilityFlags }).(pulumi.StringArrayOutput)
}

func (o ModuleWorkerScriptOutput) D1DatabaseBindings() D1DatabaseBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) D1DatabaseBindingArrayOutput { return v.D1DatabaseBindings }).(D1DatabaseBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) Dir() pulumi.StringOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringOutput { return v.Dir }).(pulumi.StringOutput)
}

func (o ModuleWorkerScriptOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

func (o ModuleWorkerScriptOutput) KvNamespaceBindings() KVNamespaceBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) KVNamespaceBindingArrayOutput { return v.KvNamespaceBindings }).(KVNamespaceBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) Logpush() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.BoolPtrOutput { return v.Logpush }).(pulumi.BoolPtrOutput)
}

func (o ModuleWorkerScriptOutput) MainModule() pulumi.StringOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringOutput { return v.MainModule }).(pulumi.StringOutput)
}

func (o ModuleWorkerScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ModuleWorkerScriptOutput) PlainTextBindings() PlainTextBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) PlainTextBindingArrayOutput { return v.PlainTextBindings }).(PlainTextBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) QueueBindings() QueueBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) QueueBindingArrayOutput { return v.QueueBindings }).(QueueBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) R2BucketBindings() R2BucketBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) R2BucketBindingArrayOutput { return v.R2BucketBindings }).(R2BucketBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) SecretTextBindings() SecretTextBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) SecretTextBindingArrayOutput { return v.SecretTextBindings }).(SecretTextBindingArrayOutput)
}

func (o ModuleWorkerScriptOutput) ServiceBindings() ServiceBindingArrayOutput {
	return o.ApplyT(func(v *ModuleWorkerScript) ServiceBindingArrayOutput { return v.ServiceBindings }).(ServiceBindingArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleWorkerScriptInput)(nil)).Elem(), &ModuleWorkerScript{})
	pulumi.RegisterOutputType(ModuleWorkerScriptOutput{})
}
