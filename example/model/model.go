package main

import (
	"context"
	"fmt"
	v1 "gitlab.inspir.work/HanYiming/cloud-go-sdk/api/v1"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/cloudgo/service"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/rest"
	"net/http"
	"time"
)

var client *service.CloudClient

func main() {
	//cloudConfig := "/Users/haminyg/Desktop/cloud/cloud-go-sdk/cloud.yaml"
	var err error
	config := &rest.Config{
		Host:    "http://11.1.2.75:31123",
		APIPath: "",
		ContentConfig: rest.ContentConfig{
			GroupVersion: "v1/cloud",
		},
		Timeout:       4 * time.Second,
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
	}

	client, err = service.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	ModelCatalog()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func ModelDownload(w http.ResponseWriter, r *http.Request) {
	c := context.Background()

	data, err := client.ApiV1().Model().Download(c, v1.ModelDownloadArgs{
		Catalog:   "1",
		ModelName: "1",
	})
	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)
}
func ModelCatalog() {
	c := context.Background()

	args := v1.ModelCatalogsArgs{
		MetaPage: v1.MetaPage{
			Page:     1,
			PageSize: 10,
		},
		MetaUuid: v1.MetaUuid{
			Uuid: "mg5rr0u8h9vhq5l",
		},
	}

	reply, err := client.ApiV1().Model().Catalogs(c, args)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range reply.List {
		fmt.Println(k, v)
	}
}
