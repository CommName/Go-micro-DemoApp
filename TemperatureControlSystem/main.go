package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
	"net/http"

	handler "github.com/CommName/Go-micro-DemoApp/TemperatureControlSystem/handler"
	subscriber "github.com/CommName/Go-micro-DemoApp/TemperatureControlSystem/subscriber"
	
	"time"
)



func TopicServer(RoomsWithThermometar *map[string] time.Time, RoomsWithAirConditioner *map[string]time.Time,internalService *micro.Service){

	RoomsWithThermometarChannel := make(chan string, 100)
	RoomsWithAirConditionerChannel := make(chan string, 100)

	
	(*internalService).Init()
	subscriber.RoomMapper(RoomsWithThermometarChannel, "Thermometar", (*internalService).Server())
	subscriber.RoomMapper(RoomsWithAirConditionerChannel, "AirConditioner", (*internalService).Server())

	go subscriber.RoomMaintainer(RoomsWithAirConditioner,RoomsWithAirConditionerChannel)
	go subscriber.RoomMaintainer(RoomsWithThermometar, RoomsWithThermometarChannel)

	if err := (*internalService).Run(); err != nil {
		log.Fatal(err)
	}

	
	close(RoomsWithThermometarChannel)
	close(RoomsWithAirConditionerChannel)
}

func main() {
	var RoomsWithThermometar map[string] time.Time
	var RoomsWithAirConditioner map[string]time.Time

	RoomsWithAirConditioner = make(map[string]time.Time)
	RoomsWithThermometar = make(map[string]time.Time)

	service := web.NewService(
		web.Name("iots.TemperatureControlSystem"),
	)

	handler := new(handler.TemperatureControlSystem)
	handler.RoomsWithThermometar = &RoomsWithThermometar
	handler.RoomsWithAirConditioner = &RoomsWithAirConditioner
	
	internalservice := micro.NewService(
		micro.Name("iots.temperature.srv.Controller"),
	)
	
	handler.InternalService = &internalservice

	service.Handle("/", http.FileServer(http.Dir("WebPages")))
	service.HandleFunc("/GetRooms", handler.GetRooms)
	service.HandleFunc("/Airconditioner/",handler.AirConditionerPage)
	service.HandleFunc("/Thermometar/",handler.ThermometarPage)
	service.HandleFunc("/SetAirconditioner/",handler.SetAirconditioner)
	

	if err := service.Init(); err!=nil {
		log.Fatal(err)
	}

	go TopicServer(&RoomsWithThermometar,&RoomsWithAirConditioner, handler.InternalService)

	if err := service.Run(); err!=nil {
		log.Fatal(err)
	}


}
