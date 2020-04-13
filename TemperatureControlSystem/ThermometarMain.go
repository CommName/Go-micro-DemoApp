package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"TemperatureControlSystem/handler"
	"TemperatureControlSystem/subscriber"

	thermometar "TemperatureControlSystem/proto/Thermometar"
)

import (
	"fmt"
)

func main() {
	var roomName string
	fmt.Print("Please input room name: ")
	fmt.Scanf("%s", &roomName)


	// New Service
	service := micro.NewService(
		micro.Name("iots.temperature.srv.Thermometar."+roomName),
		micro.Version("v1.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	termometar := new(handler.Thermometar)
	termometar.RoomName = roomName
	termometar.Temperature = 22
	thermometar.RegisterThermometarHandler(service.Server(), termometar)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iots.temperature.srv.Thermometar."+roomName, service.Server(), new(subscriber.Thermometar))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iots.temperature.srv.Thermometar."+roomName, service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
