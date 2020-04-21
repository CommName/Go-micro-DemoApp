package subscriber

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"

	TemperatureControlSystem "github.com/CommName/Go-micro-DemoApp/TemperatureControlSystem/proto/TemperatureControlSystem"
	"time"
)

func RoomMapper( roomMapper chan string, topic string, server server.Server) error{

	micro.RegisterSubscriber("iots.temperature.srv.Controller."+topic,server,func(ctx context.Context, event *TemperatureControlSystem.Message) error{
		roomMapper <- event.Say
		return nil
	})
	return nil
}

func RoomMaintainer(rooms *map[string]time.Time, channel chan string){

	deleteTimer := time.Tick (3 * time.Second)
	for ;true; {
		select {
			case roomName := <- channel:

				(*rooms)[roomName] = time.Now()
			
			case <-deleteTimer:
				for key, value := range *rooms {
					if(time.Since(value) >  3 *time.Second){
						delete(*rooms,key)
					}
				}
			
			default:
				time.Sleep(300 * time.Millisecond)
		}
	}

}