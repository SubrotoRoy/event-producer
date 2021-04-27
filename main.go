package main

import (
	"crypto/subtle"
	"os"

	"github.com/SubrotoRoy/event-producer/handler"
	"github.com/SubrotoRoy/event-producer/kafkaservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	kafkaService := kafkaservice.NewKafkaService()
	api := handler.NewEventHandler(kafkaService)
	e := echo.New()

	g := e.Group("/api/v1")

	//Added basic authentication
	g.Use(middleware.BasicAuth(BasicAuthValidator))

	g.POST("/event", api.PublishEvent)
	e.Logger.Fatal(e.Start(":8091"))
}

//BasicAuthValidator is for Basic authentication
func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {

	if subtle.ConstantTimeCompare([]byte(username), []byte(os.Getenv("USERNAME"))) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(os.Getenv("PASSWORD"))) == 1 {
		return true, nil
	}
	return false, nil
}
