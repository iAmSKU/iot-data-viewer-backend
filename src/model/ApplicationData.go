package model

type Payload struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var MessageChannel chan Payload
