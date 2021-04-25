package main

import (
	"github.com/SubrotoRoy/event-producer/handler"
	"github.com/SubrotoRoy/event-producer/kafkaservice"
	"github.com/labstack/echo"
)

func main() {
	kafkaService := kafkaservice.NewKafkaService()
	api := handler.NewEventHandler(kafkaService)
	e := echo.New()

	g := e.Group("/api/v1")

	g.POST("/event", api.PublishEvent)
	e.Logger.Fatal(e.Start(":8091"))
}
