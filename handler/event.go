package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/SubrotoRoy/event-producer/kafkaservice"
	"github.com/SubrotoRoy/event-producer/model"
	"github.com/labstack/echo"
)

type EventHandler struct {
	Kafka kafkaservice.Services
}

func NewEventHandler(kafka kafkaservice.Services) *EventHandler {
	return &EventHandler{Kafka: kafka}
}

func (e *EventHandler) PublishEvent(ctx echo.Context) error {
	event := model.Event{}

	err := json.NewDecoder(ctx.Request().Body).Decode(&event)

	if err != nil {
		errMsg := fmt.Sprint("Unable to decode request body")
		log.Println("ERROR: ", errMsg, err)
		return ctx.JSON(400, ResponseManager(nil, errors.New(errMsg)))
	}

	if event.City == "" {
		errMsg := "City name not provided"
		log.Println("ERROR: ", errMsg)
		return ctx.JSON(400, ResponseManager(nil, errors.New(errMsg)))
	}

	c := context.Background()
	err = e.Kafka.WriteToKafka(c, event)
	if err != nil {
		log.Println("Unable to post to kafka, ERROR:", err)
		return ctx.JSON(500, ResponseManager(nil, errors.New("Unable to post to kafka")))
	}

	log.Println("Saved to kafka")
	return ctx.JSON(201, ResponseManager("Event published", nil))
}

//ResponseManager creates the outgoing response
func ResponseManager(response interface{}, err error) model.APIResponse {

	apiResponse := model.APIResponse{}

	apiResponse.Response = response
	if err != nil {
		apiResponse.Error = err.Error()
	}
	return apiResponse
}
