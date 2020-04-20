package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	AirCondtioner "AirCondtioner/proto/AirCondtioner"
)

type AirCondtioner struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *AirCondtioner) Call(ctx context.Context, req *AirCondtioner.Request, rsp *AirCondtioner.Response) error {
	log.Log("Received AirCondtioner.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *AirCondtioner) Stream(ctx context.Context, req *AirCondtioner.StreamingRequest, stream AirCondtioner.AirCondtioner_StreamStream) error {
	log.Logf("Received AirCondtioner.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&AirCondtioner.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *AirCondtioner) PingPong(ctx context.Context, stream AirCondtioner.AirCondtioner_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&AirCondtioner.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
