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
	"math/rand"
	"time"
)

func EnviermentSimulator(termometar *handler.Thermometar) error {
	for ; true ; {
		temperatureChange := int64(rand.Uint64()%20) -10
		termometar.Temperature += temperatureChange
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	var roomName string
	fmt.Print("Please input room name: ")
	fmt.Scanf("%s", &roomName)


	// New Service
	service := micro.NewService(
		micro.Name("iots.temperature.srv.Thermometar."+roomName),
		micro.Version("v1.1"),
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

	//Env Simulator
	go EnviermentSimulator(termometar)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
