package internal

import (
	model "iot-data-viewer-backend/src/model"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var ConnObjectMap = make(map[string]mqtt.Client)

func AddModifyConnection(connConfig model.ConnectionConfig) {
	log.Print("Inside AddModifyConnection.")

	//Check if client already exists in the map then disconnect and again configure
	client, retStatus := ConnObjectMap[connConfig.ConnectionName]
	if retStatus {
		log.Println("Client " + connConfig.ConnectionName + " already exists, disconnecting previous connection & removing from map.")
		client.Disconnect(250)
		delete(ConnObjectMap, connConfig.ConnectionName)
	}

	retStatus, client = ConfigureAndConnect(connConfig)

	if retStatus {
		if Influx_ConfigureAndConnect(connConfig.ConnectionName) {
			if Subscribe(client, connConfig.Topic, 1) {
				ConnObjectMap[connConfig.ConnectionName] = client
				log.Println("Client " + connConfig.ConnectionName + " added in the map.")
			}
		}
	}
}
