package handler

import (
	"context"
	"sync"
	"math/rand"
	"math"
	"time"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	airconditioner "github.com/CommName/Go-micro-DemoApp/AirConditioner/proto/AirConditioner"
	Thermometar "github.com/CommName/Go-micro-DemoApp/Thermometar/proto/Thermometar"
)

type AirConditioner struct {
	RoomName string
	PowerOn bool
	HeatingMode bool
	Power int32
	Auto bool
	DesiredTemp int64
	LastTemp int64
	Mux sync.Mutex
} 



func (e *AirConditioner) SetDeviceStatus(context context.Context, req *airconditioner.DeviceStatus, rep *airconditioner.Empty) error {
	e.Mux.Lock()
	e.PowerOn = req.PowerOn
	e.HeatingMode = req.HeatingMode
	e.Power = req.Power
	e.Auto = req.AutoMode
	e.DesiredTemp = req.DesiredTemp
	e.Mux.Unlock()
	return nil
}

func (e *AirConditioner) GetDeviceStatus(context context.Context, req *airconditioner.Empty, rep *airconditioner.DeviceStatus) error {
	rep.PowerOn = e.PowerOn
	rep.HeatingMode = e.HeatingMode
	rep.Power = e.Power
	return nil
}

func AirconditionerLogic(airconditoner *AirConditioner, c client.Client){
	for ; true ; {
		airconditoner.Mux.Lock()
		if airconditoner.PowerOn {
			if airconditoner.Auto {
				if airconditoner.LastTemp > airconditoner.DesiredTemp  {
					airconditoner.HeatingMode = false
				} else {
					airconditoner.HeatingMode = true
				}
				tempDiff := int64(math.Abs(float64(airconditoner.LastTemp - airconditoner.DesiredTemp)))
				switch {
				case tempDiff >15:
					airconditoner.Power =3
				case tempDiff > 10:
					airconditoner.Power = 2
				case tempDiff > 5:
					airconditoner.Power = 1
				default:
					airconditoner.Power = 0
				}
			}
			temperatureD := (rand.Int63()%6)*int64(airconditoner.Power)
			var command string
			if airconditoner.HeatingMode{
				command = "Thermometar.CoolTheRoom"
			} else {
				command = "Thermometar.HeatTheRoom"
			}
			req := c.NewRequest("iots.temperature.srv.Thermometar."+airconditoner.RoomName, command, &Thermometar.Degrees {
				Temperature: temperatureD,
			})

			rsp := &Thermometar.RoomTemperatrue {}

			if err:= c.Call(context.TODO(),req,rsp); err!= nil {
				//Servis trenutno ne radi najverovatnija nije aktivan
				log.Log(err)

			} else {
				airconditoner.LastTemp = rsp.Temperature
			}
		}
		airconditoner.Mux.Unlock()
		time.Sleep(1500 *time.Millisecond)
	}
}