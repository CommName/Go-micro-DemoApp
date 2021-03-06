package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/CommName/Go-micro-DemoApp/Thermometar/handler"

	thermometar "github.com/CommName/Go-micro-DemoApp/Thermometar/proto/Thermometar"
	
	Controller "github.com/CommName/Go-micro-DemoApp/TemperatureControlSystem/proto/TemperatureControlSystem"
	"context"
)

import (
	"fmt"
	"time"
)


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
	termometar.AddToTemper = make(chan int64,20)
	thermometar.RegisterThermometarHandler(service.Server(), termometar)
	
	//Env Simulator
	go handler.EnviermentSimulator(termometar)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
