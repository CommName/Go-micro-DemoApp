
package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	thermometar "github.com/CommName/Go-micro-DemoApp/Thermometar/proto/Thermometar"
)

type Thermometar struct{
	RoomName string
	Temperature int64
}

func (e *Thermometar) LoadStatus(status *thermometar.RoomTemperatrue) error{
	
	status.RoomName = e.RoomName
	status.Temperature = e.Temperature
	return nil
}

func (e *Thermometar) GetStatus(context context.Context, req *thermometar.Empty, rep *thermometar.RoomTemperatrue) error {
	log.Log("Received Hello.Call request")
	err := e.LoadStatus(rep)
	return err
}

func (e *Thermometar) Subscriber(context context.Context, req *thermometar.Empty, rep *thermometar.Empty) error {
	log.Log("Received Hello.Call request")
	return nil
}

func (e *Thermometar) CoolTheRoom(context context.Context, req *thermometar.Degrees, rep *thermometar.RoomTemperatrue) error {
	log.Log("Received Hello.Call request")
	e.Temperature -= req.Temperature
	err := e.LoadStatus(rep)
	return err
}

func (e *Thermometar) HeatTheRoom(context context.Context, req *thermometar.Degrees, rep *thermometar.RoomTemperatrue) error {
	log.Log("Received Hello.Call request")
	e.Temperature += req.Temperature
	err := e.LoadStatus(rep)
	return err
}