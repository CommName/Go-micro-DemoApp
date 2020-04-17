package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"
	temperatureControlSystem "TemperatureControlSystem/proto/TemperatureControlSystem"
)

type TemperatureControlSystem struct{
	
}

// Call is a single request handler called via client.Call or the generated client code
func (e *TemperatureControlSystem) GetTemperature(ctx context.Context, req *temperatureControlSystem.Room, rsp *temperatureControlSystem.RoomTemperatrue) error {
	log.Log("Received Hello.Call request")
	return nil
}

func (e *TemperatureControlSystem) SetDesirableTemperature(ctx context.Context, req *temperatureControlSystem.RoomTemperatrue, rsp *temperatureControlSystem.Empty) error {
	log.Log("Received Hello.Call request")
	return nil
}

func (e *TemperatureControlSystem) CreateNewRoom(ctx context.Context, req *temperatureControlSystem.Room, rsp *temperatureControlSystem.Empty) error {
	log.Log("Received Hello.Call request")
	return nil
}

func (e *TemperatureControlSystem) SetUpAirConditioner(ctx context.Context, req *temperatureControlSystem.AirConditioner, rsp *temperatureControlSystem.Empty) error {
	log.Log("Received Hello.Call request")
	return nil
}

func (e *TemperatureControlSystem) DeleteRoom(ctx context.Context, req *temperatureControlSystem.Room, rsp *temperatureControlSystem.Empty) error {
	log.Log("Received Hello.Call request")
	return nil
}

