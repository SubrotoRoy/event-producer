package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/SubrotoRoy/event-producer/kafkaservice"
	"github.com/SubrotoRoy/event-producer/model"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	Kafka kafkaservice.Services
}

func NewEventHandler(kafka kafkaservice.Services) *EventHandler {
	return &EventHandler{Kafka: kafka}
}

//PublishEvent handler to handle /event POST calls
func (e *EventHandler) PublishEvent(ctx echo.Context) error {
	event := model.Event{}

	err := json.NewDecoder(ctx.Request().Body).Decode(&event)

	if err != nil {
		errmsg := "unable to decode request body"
		log.Println("ERROR: ", errmsg, err)
		return ctx.JSON(400, ResponseManager(nil, errors.New(errmsg)))
	}

	//Returning error if city name is not provided
	if event.City == "" {
		errmsg := "city name not provided"
		log.Println("ERROR: ", errmsg)
		return ctx.JSON(400, ResponseManager(nil, errors.New(errmsg)))
	}

	c := context.Background()
	err = e.Kafka.WriteToKafka(c, event)

	//Returning error if post to kafka is unsuccessful
	if err != nil {
		log.Println("Unable to post to kafka, ERROR:", err)
		return ctx.JSON(500, ResponseManager(nil, errors.New("unable to post to kafka")))
	}

	log.Println("Saved to kafka")
	return ctx.JSON(201, ResponseManager("Event published", nil))
}

//ResponseManager creates the outgoing response structure
func ResponseManager(response interface{}, err error) model.APIResponse {

	apiResponse := model.APIResponse{}

	apiResponse.Response = response
	if err != nil {
		apiResponse.Error = err.Error()
	}
	return apiResponse
}
