package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"TemperatureControlSystem/handler"
	"TemperatureControlSystem/subscriber"

	TemperatureControlSystem "TemperatureControlSystem/proto/TemperatureControlSystem"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("iots.temperature.srv.TemperatureControlSystem"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	TemperatureControlSystem.RegisterTemperatureControlSystemHandler(service.Server(), new(handler.TemperatureControlSystem))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iots.temperature.srv.TemperatureControlSystem", service.Server(), new(subscriber.TemperatureControlSystem))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iots.temperature.srv.TemperatureControlSystem", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
