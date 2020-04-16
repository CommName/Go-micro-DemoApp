package main

import (

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"TemperatureControlSystem/handler"
	"TemperatureControlSystem/subscriber"
	AirConditioner "TemperatureControlSystem/proto/AirConditioner"
	
)

import (
	"context"
	Thermometar "TemperatureControlSystem/proto/Thermometar"
	"github.com/micro/go-micro/client"
	//"github.com/micro/go-micro/metadata"
)

import (
	"fmt"
	"math/rand"
	"time"
)

func airconditionerLogic(airconditoner *handler.AirConditioner, c client.Client){
	for ; true ; {
		if airconditoner.PowerOn {
			temperatureD := rand.Int63()%3*int64(airconditoner.Power)
			var command string
			if airconditoner.HeatingMode{
				command = "Thermometar.CoolTheRoom"
			} else {
				command = "Thermometar.CoolTheRoom"
			}
			req := c.NewRequest("iots.temperature.srv.Thermometar."+airconditoner.RoomName, command, &Thermometar.Degrees {
				Temperature: temperatureD,
			})

			rsp := &Thermometar.RoomTemperatrue {}

			if err:= c.Call(context.TODO(),req,rsp); err!= nil {
				//Servis trenutno ne radi najverovatnija greska
				fmt.Println(err)

			} else {
				fmt.Printf("Room %v has temperature %v", rsp.RoomName, rsp.Temperature)
			}
		}
		time.Sleep(5000 *time.Millisecond)
	}

}

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

	service.Init()

	// Register Handler
	
	AirConditioner.RegisterAirConditionerHandler(service.Server(), airconditoner)

	// Register Struct as Subscriber
	micro.RegisterSubscriber("iots.temperature.srv.AirConditioner", service.Server(), new(subscriber.AirConditioner))

	// Register Function as Subscriber
	micro.RegisterSubscriber("iots.temperature.srv.AirConditioner", service.Server(), subscriber.Handler)
	
	go airconditionerLogic(airconditoner, service.Client())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
