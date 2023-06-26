package v1

import (
	"context"
	v1 "gitlab.inspir.work/HanYiming/cloud-go-sdk/api/v1"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/pkg/json"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/rest"
)

type LogGetter interface {
	Log() ILog
}

type ILog interface {
	Delete(c context.Context, args v1.LogDeleteArgs) error
	Download(c context.Context, args v1.LogDownloadArgs) ([]byte, error)
	List(c context.Context, args v1.LogListArgs) (*v1.LogListReply, error)
	Detail(c context.Context, args v1.LogDetailArgs) (*v1.LogDetailReply, error)
}

type log struct {
	client rest.Interface
}

func (l *log) Delete(c context.Context, args v1.LogDeleteArgs) error {
	return nil
}

func (l *log) Download(c context.Context, args v1.LogDownloadArgs) (bytes []byte, err error) {
	bytes, err = l.client.
		Get().
		Resource("log").
		SubResource("download").
		VersionedParams(args).
		Do(c).FileData()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (l *log) List(c context.Context, args v1.LogListArgs) (result *v1.LogListReply, err error) {

	data, err := l.client.
		Get().
		Resource("log").
		SubResource("list").
		VersionedParams(args).
		Do(c).
		Data()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	return
}

func (l *log) Detail(c context.Context, args v1.LogDetailArgs) (*v1.LogDetailReply, error) {
	return nil, nil
}

func newLog(c *APIV1Client) *log {
	return &log{
		client: c.RESTClient(),
	}
}
