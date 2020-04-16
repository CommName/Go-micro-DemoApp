package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	airConditioner "TemperatureControlSystem/proto/AirConditioner"
)

type AirConditioner struct {}

func (e *AirConditioner) Handle(ctx context.Context, msg *airConditioner.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}