package main

import (
	"github.com/micro/go-micro"
	"fmt"
	"context"
	hello "hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.service.hello.client"), //name the client service
	)

	// Initialise service
	service.Init()

	helloClient := hello.NewHelloService("com.foo.srv.hello", service.Client())

	resp, err := helloClient.Helloo(context.TODO(), &hello.Request { Name: "World"})

	if err!= nil {
		fmt.Print(err)
		return
	}
	
	fmt.Print(resp.Msg)

}
