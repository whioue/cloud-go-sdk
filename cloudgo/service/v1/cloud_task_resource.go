package v1

import (
	"context"
	v1 "github.com/whioue/cloud-go-sdk/api/v1"
	"github.com/whioue/cloud-go-sdk/pkg/json"
	"github.com/whioue/cloud-go-sdk/rest"
)

var _ ITaskResource = &taskResource{}

type TaskResourceGetter interface {
	TaskResource() ITaskResource
}

type ITaskResource interface {
	Start(c context.Context, args v1.TaskResourceStartArgs) (*v1.TaskResourceStartReply, error)
	Restart(c context.Context, args v1.TaskResourceRestartArgs) error
	Stop(c context.Context, args v1.TaskResourceStopArgs) error
	Detail(c context.Context, args v1.TaskResourceDetailArgs) (*v1.TaskResourceDetailReply, error)
	TensorboardDetail(c context.Context, args v1.TaskResourceTensorboardDetailArgs) (*v1.TaskResourceTensorboardDetailReply, error)
}

type taskResource struct {
	client rest.Interface
}

func newTaskResource(c *APIV1Client) *taskResource {
	return &taskResource{
		client: c.RESTClient(),
	}
}

func (t *taskResource) Start(c context.Context, args v1.TaskResourceStartArgs) (result *v1.TaskResourceStartReply, err error) {
	data, err := t.client.
		Post().
		Resource("task").
		SubResource("resource", "start").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func (t *taskResource) Restart(c context.Context, args v1.TaskResourceRestartArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("resource", "restart").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t *taskResource) Stop(c context.Context, args v1.TaskResourceStopArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("resource", "stop").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t *taskResource) Detail(c context.Context, args v1.TaskResourceDetailArgs) (result *v1.TaskResourceDetailReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("resource").
		VersionedParams(args).
		Do(c).
		Data()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func (t *taskResource) TensorboardDetail(c context.Context, args v1.TaskResourceTensorboardDetailArgs) (result *v1.TaskResourceTensorboardDetailReply, err error) {
	data, err := t.client.
		Post().
		Resource("task").
		SubResource("resource", "tensorboard").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}
