package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	temperatureControlSystem "TemperatureControlSystem/proto/TemperatureControlSystem"

)

type TemperatureControlSystem struct{}

func (e *TemperatureControlSystem) Handle(ctx context.Context, msg *temperatureControlSystem.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *temperatureControlSystem.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

