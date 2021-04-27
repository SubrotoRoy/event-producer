package model

//APIResponse struct is to send out reponse
type APIResponse struct {
	Response interface{} `json:"response"`
	Error    string      `json:"error"`
}

//Event struct is to accept request body and structure to send messages to kafka
type Event struct {
	FuelLid bool   `json:"fuellid"`
	City    string `json:"city"`
}
