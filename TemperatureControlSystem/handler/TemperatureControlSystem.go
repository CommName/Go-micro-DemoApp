package handler

import (

	"net/http"
	"context"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"encoding/json"
	"time"
	"strings"
	Thermometar "github.com/CommName/Go-micro-DemoApp/Thermometar/proto/Thermometar"
	Airconditioner "github.com/CommName/Go-micro-DemoApp/AirConditioner/proto/AirConditioner"

)

type TemperatureControlSystem struct{
	RoomsWithThermometar *map[string] time.Time
	RoomsWithAirConditioner *map[string]time.Time
	InternalService *micro.Service
}


func (t *TemperatureControlSystem)GetRooms(w http.ResponseWriter, r *http.Request){
	var RoomNames []string
	for key, _ := range *t.RoomsWithThermometar {
		RoomNames = append(RoomNames,key)
	}
	
	//TODO add other rooms

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(RoomNames); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (t *TemperatureControlSystem)ThermometarPage(w http.ResponseWriter, r *http.Request){
	RoomName :=strings.TrimPrefix( r.URL.RequestURI(),"/Thermometar/")
	if _, exists := (*t.RoomsWithThermometar)[RoomName]; exists {
		c := (*t.InternalService).Client()
		req := c.NewRequest("iots.temperature.srv.Thermometar."+RoomName, "Thermometar.GetStatus", &Thermometar.Empty {
	
		})
		rsp := &Thermometar.RoomTemperatrue {}
		if err:= c.Call(context.TODO(),req,rsp); err!= nil {
			http.Error(w, err.Error(), 500)
			log.Log(err)
			return
		}
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		response := map[string]interface{}{
			"RoomName": RoomName,
			"Temperature": "No data available",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}

func (t *TemperatureControlSystem)AirConditionerPage(w http.ResponseWriter, r *http.Request){
	RoomName :=strings.TrimPrefix( r.URL.RequestURI(),"/Airconditioner/")

	if _, exists := (*t.RoomsWithAirConditioner)[RoomName]; exists {
		c := (*t.InternalService).Client()
		req := c.NewRequest("iots.temperature.srv.AirConditioner."+RoomName, "AirConditioner.GetDeviceStatus", &Airconditioner.Empty {
	
		})
	
		rsp := &Airconditioner.DeviceStatus {}
		if err:= c.Call(context.TODO(),req,rsp); err!= nil {
			http.Error(w, err.Error(), 500)
			log.Log(err)
			return
		}
	
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	} else {
		response := map[string]interface{}{
			"powerOn": "No data available",
			"HeatingMode": "No data available",
			"Power": "No data available",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	
}

func (t *TemperatureControlSystem)SetAirconditioner(w http.ResponseWriter, r*http.Request){
	RoomName :=strings.TrimPrefix( r.URL.RequestURI(),"/SetAirconditioner/")
	log.Log("Test")
	if _, exists := (*t.RoomsWithAirConditioner)[RoomName]; exists {
		// decode the incoming request as json
		var request map[string]interface{}
		log.Log("test")
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		
		c := (*t.InternalService).Client()
		req := c.NewRequest("iots.temperature.srv.AirConditioner."+RoomName, "AirConditioner.SetDeviceStatus", &Airconditioner.DeviceStatus {
			PowerOn: request["PowerOn"].(bool),
			HeatingMode: request["heatingOn"].(bool),
			Power: int32(request["Power"].(float64)),
		})
	
		rsp := &Airconditioner.Empty {}
		if err:= c.Call(context.TODO(),req,rsp); err!= nil {
			http.Error(w, err.Error(), 500)
			log.Log(err)
			return
		}
	
		// encode and write the response as json
		if err := json.NewEncoder(w).Encode(rsp); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
}