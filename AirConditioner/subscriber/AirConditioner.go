package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	AirConditioner "AirConditioner/proto/AirConditioner"
)

type AirConditioner struct{}

func (e *AirConditioner) Handle(ctx context.Context, msg *AirConditioner.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *AirConditioner.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
