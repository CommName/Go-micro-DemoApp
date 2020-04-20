package handler

import (
	"context"
	
	airconditioner "github.com/CommName/Go-micro-DemoApp/AirConditioner/proto/AirConditioner"
)

type AirConditioner struct {
	RoomName string
	PowerOn bool
	HeatingMode bool
	Power int32
} 



func (e *AirConditioner) SetDeviceStatus(context context.Context, req *airconditioner.DeviceStatus, rep *airconditioner.Empty) error {
	e.PowerOn = req.PowerOn
	e.HeatingMode = req.HeatingMode
	e.Power = req.Power
	return nil
}

func (e *AirConditioner) GetDeviceStatus(context context.Context, req *airconditioner.Empty, rep *airconditioner.DeviceStatus) error {
	rep.PowerOn = e.PowerOn
	rep.HeatingMode = e.HeatingMode
	rep.Power = e.Power
	return nil
}