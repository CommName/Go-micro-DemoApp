package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/server"
	"context"
	"net/http"


	TemperatureControlSystem "TemperatureControlSystem/proto/TemperatureControlSystem"
	"fmt"
	"time"
	"strings"
	"strconv"
	"encoding/json"
)

var (
	RoomsWithThermometar map[string] time.Time
	RoomsWithAirConditioner map[string]time.Time
	internalService micro.Service
)

func RoomMapper( roomMapper chan string, topic string, server server.Server) error{

	micro.RegisterSubscriber("iots.temperature.srv.Controller."+topic,server,func(ctx context.Context, event *TemperatureControlSystem.Message) error{
		roomMapper <- event.Say
		return nil
	})
	return nil
}

func RoomMaintainer(rooms *map[string]time.Time, channel chan string){

	deleteTimer := time.Tick (3 * time.Second)
	for ;true; {
		select {
			case roomName := <- channel:

				(*rooms)[roomName] = time.Now()
			
			case <-deleteTimer:
				for key, value := range *rooms {
					if(time.Since(value) >  3 *time.Second){
						delete(*rooms,key)
					}
				}
			
			default:
				time.Sleep(300 * time.Millisecond)
		}
	}

}

func AirConditionerPage(w http.ResponseWriter, r *http.Request){
	RoomName :=strings.TrimPrefix( r.URL.RequestURI(),"/Airconditioner/")
	
}

func ThermometarPage(w http.ResponseWriter, r *http.Request){
	RoomName :=strings.TrimPrefix( r.URL.RequestURI(),"/Thermometar/")
	
}

func GetRooms(w http.ResponseWriter, r *http.Request){
	var RoomNames []string
	for key, _ := range RoomsWithThermometar {
		RoomNames = append(RoomNames,key)
	}
	
	//TODO add other rooms

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(RoomNames); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TopicServer(){
	
	RoomsWithThermometarChannel := make(chan string, 100)
	RoomsWithAirConditionerChannel := make(chan string, 100)

	internalService = micro.NewService(
		micro.Name("iots.temperature.srv.Controller"),
	)
	internalService.Init()
	RoomMapper(RoomsWithThermometarChannel, "Thermometar", internalService.Server())
	RoomMapper(RoomsWithAirConditionerChannel, "AirConditioner", internalService.Server())

	go RoomMaintainer(&RoomsWithAirConditioner,RoomsWithAirConditionerChannel)
	go RoomMaintainer(&RoomsWithThermometar, RoomsWithThermometarChannel)

	if err := internalService.Run(); err != nil {
		log.Fatal(err)
	}

	
	close(RoomsWithThermometarChannel)
	close(RoomsWithAirConditionerChannel)
}

func main() {
	RoomsWithAirConditioner = make(map[string]time.Time)
	RoomsWithThermometar = make(map[string]time.Time)

	service := web.NewService(
		web.Name("iots.TemperatureControlSystem"),
	)

	service.Handle("/", http.FileServer(http.Dir("WebPages")))
	service.HandleFunc("/GetRooms", GetRooms)
	service.HandleFunc("/Airconditioner/",AirConditionerPage)
	service.HandleFunc("/Thermometar/",ThermometarPage)
	

	if err := service.Init(); err!=nil {
		log.Fatal(err)
	}

	go TopicServer()

	if err := service.Run(); err!=nil {
		log.Fatal(err)
	}


}
