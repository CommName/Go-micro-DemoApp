package main

import (

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/CommName/Go-micro-DemoApp/AirConditioner/handler"
	AirConditioner "github.com/CommName/Go-micro-DemoApp/AirConditioner/proto/AirConditioner"
	
)

import (
	"context"
	Thermometar "github.com/CommName/Go-micro-DemoApp/Thermometar/proto/Thermometar"
	"github.com/micro/go-micro/client"
	Controller "github.com/CommName/Go-micro-DemoApp/TemperatureControlSystem/proto/TemperatureControlSystem"
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
				//Servis trenutno ne radi najverovatnija nije aktivan
				log.Log(err)

			} else {
				fmt.Printf("Room %v has temperature %v\n", rsp.RoomName, rsp.Temperature)
			}
		}
		time.Sleep(1500 *time.Millisecond)
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

	service.Init(
		//Publisher logic to broadcast room name
		micro.AfterStart(func () error {
			topic := micro.NewPublisher("iots.temperature.srv.Controller.AirConditioner", service.Client())

			message := &Controller.Message{
					Say: airconditoner.RoomName,
			}
			fmt.Println("Poruka kreirana")
		
			for ; true ; {
				err := topic.Publish(context.TODO(), message)
				if (err!=nil){
					log.Fatal(err)
				}
				
				fmt.Println("Poruka poslata")
				time.Sleep(3 * time.Second)
			}
		
			return nil
		}),
	)

	// Register Handler
	AirConditioner.RegisterAirConditionerHandler(service.Server(), airconditoner)

	go airconditionerLogic(airconditoner, service.Client())

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
