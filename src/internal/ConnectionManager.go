package internal

import (
	model "iot-data-viewer-backend/src/model"
	"log"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ConnectionObject struct {
	Mu         sync.RWMutex
	ConnObject map[string]mqtt.Client
}

var connectionObjectMap ConnectionObject

func AddModifyConnection(connConfig model.ConnectionConfig) {
	log.Print("Inside AddModifyConnection...")

	var connectivity model.MqttConfiguration
	connectivity.ClientName = connConfig.ConnectionName
	connectivity.HostName = connConfig.HostName
	connectivity.HostPort = connConfig.HostPort
	connectivity.UserName = connConfig.UserName
	connectivity.Password = connConfig.Password
	connectivity.Topic = connConfig.Topic

	ret, opts := connectivity.Configure()

	if ret {
		ret, client := connectivity.Connect(opts)
		if ret {
			if connectivity.Subscribe(client, connConfig.Topic, 1) {
				connectionObjectMap.ConnObject[connConfig.ConnectionName] = client
			}
		}
	}
}
