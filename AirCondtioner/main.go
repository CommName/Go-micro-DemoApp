package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"AirCondtioner/handler"
	"AirCondtioner/subscriber"

	AirCondtioner "AirCondtioner/proto/AirCondtioner"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("github.com/CommName.srv.AirCondtioner"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	AirCondtioner.RegisterAirCondtionerHandler(service.Server(), new(handler.AirCondtioner))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("github.com/CommName.srv.AirCondtioner", service.Server(), new(subscriber.AirCondtioner))

	// Register Function as Subscriber
	micro.RegisterSubscriber("github.com/CommName.srv.AirCondtioner", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
