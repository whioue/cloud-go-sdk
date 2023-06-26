package main

import (
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/cloudgo/service"
	"gitlab.inspir.work/HanYiming/cloud-go-sdk/rest"
	"time"
)

var client *service.CloudClient

func main() {
	//cloudConfig := "/Users/haminyg/Desktop/cloud/cloud-go-sdk/cloud.yaml"
	var err error
	config := &rest.Config{
		Host:    "http://127.0.0.1:8100",
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

}
