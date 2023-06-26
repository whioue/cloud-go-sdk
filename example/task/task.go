package main

import (
	"context"
	"fmt"
	v1 "github.com/whioue/cloud-go-sdk/api/v1"
	"github.com/whioue/cloud-go-sdk/cloudgo/service"
	"github.com/whioue/cloud-go-sdk/rest"
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

	List()
}

func Create() {
	c := context.Background()

	args := v1.TaskCreateArgs{}

	reply, err := client.ApiV1().Task().Create(c, args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}

func List() {
	c := context.Background()

	args := v1.TaskListArgs{
		Page:     0,
		PageSize: 0,
		Status:   "",
		Name:     "",
	}

	reply, err := client.ApiV1().Task().List(c, args)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range reply.List {
		fmt.Println(k, v.Name)
	}
}

func Delete() {
	c := context.Background()

	args := v1.TaskDeleteArgs{}

	err := client.ApiV1().Task().Delete(c, args)
	if err != nil {
		fmt.Println(err)
	}
}

func Detail() {
	c := context.Background()

	args := v1.TaskDetailArgs{}

	reply, err := client.ApiV1().Task().Detail(c, args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
}
