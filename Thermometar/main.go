package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"Thermometar/handler"

	thermometar "Thermometar/proto/Thermometar"
	
	Controller "TemperatureControlSystem/proto/TemperatureControlSystem"
	"context"
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
	service.Init(
		//Publisher logic to broadcast room name
		micro.AfterStart(func () error {
			topic := micro.NewPublisher("iots.temperature.srv.Controller.Thermometar", service.Client())

			message := &Controller.Message{
					Say: roomName,
			}
		
			for ; true ; {
				err := topic.Publish(context.TODO(), message)
				if (err!=nil){
					log.Fatal(err)
				}
				time.Sleep(3 * time.Second)
			}
		
			return nil
		}),
	)

	// Register Handler
	termometar := new(handler.Thermometar)
	termometar.RoomName = roomName
	termometar.Temperature = 22
	thermometar.RegisterThermometarHandler(service.Server(), termometar)
	
	//Env Simulator
	go EnviermentSimulator(termometar)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
