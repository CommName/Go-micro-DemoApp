package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	AirCondtioner "AirCondtioner/proto/AirCondtioner"
)

type AirCondtioner struct{}

func (e *AirCondtioner) Handle(ctx context.Context, msg *AirCondtioner.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *AirCondtioner.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
