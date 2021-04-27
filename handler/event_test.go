package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/SubrotoRoy/event-producer/handler"
	"github.com/SubrotoRoy/event-producer/kafkaservice"
	"github.com/SubrotoRoy/event-producer/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var mockKafkaSvc *kafkaservice.MockKafkaSvc

var api *handler.EventHandler

func init() {
	mockKafkaSvc = &kafkaservice.MockKafkaSvc{}
	api = handler.NewEventHandler(mockKafkaSvc)
	os.Setenv("USERNAME", "admin")
	os.Setenv("PASSWORD", "password")
}

func TestResponseManager(t *testing.T) {
	someMessage := "some message"
	actual := handler.ResponseManager(someMessage, nil)

	assert.Equal(t, someMessage, actual.Response)
}

func TestResponseManagerWithError(t *testing.T) {
	someMessage := "some message"
	actual := handler.ResponseManager(someMessage, errors.New(someMessage))

	assert.Equal(t, someMessage, actual.Error)
}

func TestPublishEventNoData(t *testing.T) {
	var inputJSON = ``

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", strings.NewReader(inputJSON))
	req.Header.Set("Content-Type", "applicaiton/json")
	req.SetBasicAuth("admin", "password")
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	if assert.NoError(t, api.PublishEvent(context)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestPublishEventNoCity(t *testing.T) {
	var inputJSON = `{
			"fuellid":true
		}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", strings.NewReader(inputJSON))
	req.Header.Set("Content-Type", "applicaiton/json")
	req.SetBasicAuth("admin", "password")
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	if assert.NoError(t, api.PublishEvent(context)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestPublishEvent(t *testing.T) {
	var inputJSON = `{
			"fuellid":true,
			"city":"ranchi"
		}`

	event := model.Event{}
	json.Unmarshal([]byte(inputJSON), &event)
	mockKafkaSvc.On("WriteToKafka", context.Background(), event).Return(nil)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", strings.NewReader(inputJSON))
	req.Header.Set("Content-Type", "applicaiton/json")
	req.SetBasicAuth("admin", "password")
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	if assert.NoError(t, api.PublishEvent(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestPublishEventKafkaFailure(t *testing.T) {
	var inputJSON = `{
			"fuellid":true,
			"city":"jaipur"
		}`

	event := model.Event{}
	json.Unmarshal([]byte(inputJSON), &event)
	mockKafkaSvc.On("WriteToKafka", context.Background(), event).Return(errors.New("some error"))
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", strings.NewReader(inputJSON))
	req.Header.Set("Content-Type", "applicaiton/json")
	req.SetBasicAuth("admin", "password")
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	if assert.NoError(t, api.PublishEvent(context)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}
