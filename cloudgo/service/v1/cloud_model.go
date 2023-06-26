package v1

import (
	"context"
	v1 "gitlab.inspir.work/HanYiming/cloud-go-sdk/api/v1"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/pkg/json"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/rest"
)

type ModelGetter interface {
	Model() IModel
}

type IModel interface {
	Delete(c context.Context, args v1.ModelDeleteArgs) error
	Catalogs(c context.Context, args v1.ModelCatalogsArgs) (*v1.ModelCatalogsReply, error)
	Download(c context.Context, args v1.ModelDownloadArgs) ([]byte, error)
	List(c context.Context, args v1.ModelListArgs) (*v1.ModelListReply, error)
}

type model struct {
	client rest.Interface
}

func newModel(c *APIV1Client) *model {
	return &model{
		client: c.RESTClient(),
	}
}

func (m *model) Delete(c context.Context, args v1.ModelDeleteArgs) error {
	return nil
}

func (m *model) Catalogs(c context.Context, args v1.ModelCatalogsArgs) (result *v1.ModelCatalogsReply, err error) {
	data, err := m.client.
		Get().
		Resource("task").
		SubResource("model", "catalogs").
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

func (m *model) Download(c context.Context, args v1.ModelDownloadArgs) (bytes []byte, err error) {
	bytes, err = m.client.
		Get().
		Resource("task").
		SubResource("model", "download").
		VersionedParams(args).
		Do(c).FileData()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (m *model) List(c context.Context, args v1.ModelListArgs) (result *v1.ModelListReply, err error) {
	data, err := m.client.
		Get().
		Resource("task").
		SubResource("model", "list").
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
