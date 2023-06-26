package v1

import (
	"context"
	"encoding/json"
	v1 "github.com/whioue/cloud-go-sdk/api/v1"
	"github.com/whioue/cloud-go-sdk/rest"
)

var _ ITask = &task{}

type TaskGetter interface {
	Task() ITask
}

type ITask interface {
	Create(c context.Context, args v1.TaskCreateArgs) (*v1.TaskCreateReply, error)
	Delete(c context.Context, args v1.TaskDeleteArgs) error
	Detail(c context.Context, args v1.TaskDetailArgs) (*v1.TaskDetailReply, error)
	List(c context.Context, args v1.TaskListArgs) (*v1.TaskListReply, error)
	Update(c context.Context, args v1.TaskUpdateArgs) error

	Start(c context.Context, args v1.TaskStartArgs) error
	Stop(c context.Context, args v1.TaskStopArgs) error
	PodList(c context.Context, args v1.TaskPodListArgs) (*v1.TaskPodListReply, error)
}

type task struct {
	client rest.Interface
}

func (t *task) Start(c context.Context, args v1.TaskStartArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("start").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t *task) Stop(c context.Context, args v1.TaskStopArgs) error {
	_, err := t.client.
		Post().
		Resource("task").
		SubResource("stop").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}

func (t *task) PodList(c context.Context, args v1.TaskPodListArgs) (result *v1.TaskPodListReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("pods").
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

func newTask(c *APIV1Client) *task {
	return &task{
		client: c.RESTClient(),
	}
}

func (t *task) Create(c context.Context, args v1.TaskCreateArgs) (result *v1.TaskCreateReply, err error) {
	data, err := t.client.
		Post().
		Resource("task").
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
func (t *task) Delete(c context.Context, args v1.TaskDeleteArgs) error {
	_, err := t.client.
		Delete().
		Resource("task").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}
	return nil
}
func (t *task) Detail(c context.Context, args v1.TaskDetailArgs) (result *v1.TaskDetailReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
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
func (t *task) List(c context.Context, args v1.TaskListArgs) (result *v1.TaskListReply, err error) {
	data, err := t.client.
		Get().
		Resource("task").
		SubResource("list").
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
func (t *task) Update(c context.Context, args v1.TaskUpdateArgs) error {
	_, err := t.client.
		Put().
		Resource("task").
		Body(args).
		Do(c).
		Data()
	if err != nil {
		return err
	}

	return nil
}
