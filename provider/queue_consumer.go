package provider

import (
	"encoding/json"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/queues"
	provider "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type (
	QueueConsumer struct{}

	QueueConsumerArgs struct {
		AccountID              string `pulumi:"accountId"`
		QueueID                string `pulumi:"queueId"`
		ScriptName             string `pulumi:"scriptName"`
		DeadLetterQueue        string `pulumi:"deadLetterQueue,optional"`
		BatchSize              *int   `pulumi:"batchSize,optional"`
		MaxRetries             *int   `pulumi:"maxRetries,optional"`
		MaxWaitTimeMs          *int   `pulumi:"maxWaitTimeMs,optional"`
		RetryDelay             *int   `pulumi:"retryDelay,optional"`
		MaxConsumerConcurrency *int   `pulumi:"maxConsumerConcurrency,optional"`
	}
	QueueConsumerState struct {
		QueueConsumerArgs
		ConsumerID string `pulumi:"consumerId"`
	}

	queueConsumerParams struct {
		Type            string                 `json:"type"`
		ScriptName      string                 `json:"script_name"`
		DeadLetterQueue string                 `json:"dead_letter_queue,omitempty"`
		Settings        *queueConsumerSettings `json:"settings,omitempty"`
	}
	queueConsumerSettings struct {
		BatchSize      *int `json:"batch_size,omitempty"`
		MaxRetries     *int `json:"max_retries,omitempty"`
		MaxWaitTimeMs  *int `json:"max_wait_time_ms,omitempty"`
		RetryDelay     *int `json:"retry_delay,omitempty"`
		MaxConcurrency *int `json:"max_concurrency,omitempty"`
	}
)

func createQueueConsumerParams(input QueueConsumerArgs) queueConsumerParams {
	params := queueConsumerParams{
		Type:            "worker",
		ScriptName:      input.ScriptName,
		DeadLetterQueue: input.DeadLetterQueue,
	}
	if input.BatchSize != nil || input.MaxRetries != nil || input.MaxWaitTimeMs != nil || input.RetryDelay != nil || input.MaxConsumerConcurrency != nil {
		params.Settings = &queueConsumerSettings{
			BatchSize:      input.BatchSize,
			MaxRetries:     input.MaxRetries,
			MaxWaitTimeMs:  input.MaxWaitTimeMs,
			RetryDelay:     input.RetryDelay,
			MaxConcurrency: input.MaxConsumerConcurrency,
		}
	}
	return params
}

// Create implements infer.CustomResource.
func (q *QueueConsumer) Create(ctx provider.Context, name string, input QueueConsumerArgs, preview bool) (id string, output QueueConsumerState, err error) {
	output.QueueConsumerArgs = input
	if preview {
		return id, output, nil
	}

	cli := getConfig(ctx).client
	resp, err := cli.Queues.Consumers.New(ctx, input.QueueID, queues.ConsumerNewParams{
		AccountID: cloudflare.String(input.AccountID),
		Body:      createQueueConsumerParams(input),
	})
	if err != nil {
		return id, output, err
	}

	var raw struct {
		ConsumerID string `json:"consumer_id"`
	}
	err = json.Unmarshal([]byte(resp.JSON.RawJSON()), &raw)
	if err != nil {
		return id, output, err
	}
	output.ConsumerID = raw.ConsumerID

	return id, output, err
}

// Update implements infer.CustomUpdate.
func (q *QueueConsumer) Update(ctx provider.Context, id string, olds QueueConsumerState, news QueueConsumerArgs, preview bool) (QueueConsumerState, error) {
	cli := getConfig(ctx).client
	_, err := cli.Queues.Consumers.Update(ctx, olds.QueueID, olds.ConsumerID, queues.ConsumerUpdateParams{
		AccountID: cloudflare.String(olds.AccountID),
		Body:      nil,
	})
	_ = err
	panic("unimplemented")
}

// Delete implements infer.CustomDelete.
func (q *QueueConsumer) Delete(ctx provider.Context, id string, props QueueConsumerState) error {
	panic("unimplemented")
}

// Diff implements infer.CustomDiff.
func (q *QueueConsumer) Diff(ctx provider.Context, id string, olds QueueConsumerState, news QueueConsumerArgs) (provider.DiffResponse, error) {
	panic("unimplemented")
}

var _ infer.CustomResource[QueueConsumerArgs, QueueConsumerState] = (*QueueConsumer)(nil)
var _ infer.CustomUpdate[QueueConsumerArgs, QueueConsumerState] = (*QueueConsumer)(nil)
var _ infer.CustomDelete[QueueConsumerState] = (*QueueConsumer)(nil)
var _ infer.CustomDiff[QueueConsumerArgs, QueueConsumerState] = (*QueueConsumer)(nil)
