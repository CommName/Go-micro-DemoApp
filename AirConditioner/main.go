package main

import (

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/CommName/Go-micro-DemoApp/AirConditioner/handler"
	AirConditioner "github.com/CommName/Go-micro-DemoApp/AirConditioner/proto/AirConditioner"
	
)

import (
	"context"
	Controller "github.com/CommName/Go-micro-DemoApp/TemperatureControlSystem/proto/TemperatureControlSystem"
)

import (
	"fmt"
	"time"
)




func main() {
	airconditoner := new(handler.AirConditioner)
	fmt.Print("Insert the name of the room: ")
	fmt.Scanf("%s" , &airconditoner.RoomName)
	airconditoner.PowerOn = true
	airconditoner.HeatingMode = true
	airconditoner.Power = 1;

	service := micro.NewService(
		micro.Name("iots.temperature.srv.AirConditioner."+airconditoner.RoomName),
		micro.Version("v1.0"),
	)

	service.Init(
		//Publisher logic to broadcast room name
		micro.AfterStart(func () error {
			topic := micro.NewPublisher("iots.temperature.srv.Controller.AirConditioner", service.Client())

			message := &Controller.Message{
					Say: airconditoner.RoomName,
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
	AirConditioner.RegisterAirConditionerHandler(service.Server(), airconditoner)

	go handler.AirconditionerLogic(airconditoner, service.Client())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
