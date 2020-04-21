
package handler

import (
	"context"
	"math/rand"
	"time"
	thermometar "github.com/CommName/Go-micro-DemoApp/Thermometar/proto/Thermometar"
)

type Thermometar struct{
	RoomName string
	Temperature int64
	AddToTemper chan int64
}

func (e *Thermometar) LoadStatus(status *thermometar.RoomTemperatrue) error{
	
	status.RoomName = e.RoomName
	status.Temperature = e.Temperature
	return nil
}

func (e *Thermometar) GetStatus(context context.Context, req *thermometar.Empty, rep *thermometar.RoomTemperatrue) error {
	err := e.LoadStatus(rep)
	return err
}



func (e *Thermometar) CoolTheRoom(context context.Context, req *thermometar.Degrees, rep *thermometar.RoomTemperatrue) error {
	e.AddToTemper <- -req.Temperature
	err := e.LoadStatus(rep)
	return err
}

func (e *Thermometar) HeatTheRoom(context context.Context, req *thermometar.Degrees, rep *thermometar.RoomTemperatrue) error {
	e.AddToTemper <- req.Temperature
	err := e.LoadStatus(rep)
	return err
}

func EnviermentSimulator(termometar *Thermometar) error {
	for ; true ; {
		
		select {
		case diff := <- termometar.AddToTemper:
				termometar.Temperature += diff
				time.Sleep(100*time.Millisecond)

			default:

				temperatureChange := int64(rand.Uint64()%20) -10
				termometar.Temperature += temperatureChange
				time.Sleep(1000 * time.Millisecond)

		}
	}
	return nil
}