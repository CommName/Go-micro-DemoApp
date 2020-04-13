package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	thermometar "TemperatureControlSystem/proto/Thermometar"
)

type Thermometar struct {}

func (e *Thermometar) Handle(ctx context.Context, msg *thermometar.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler2(ctx context.Context, msg *thermometar.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
