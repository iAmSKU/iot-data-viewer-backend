package internal

import (
	"fmt"
	"iot-data-viewer-backend/src/model"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConfigureAndConnect(conData model.ConnectionConfig) (bool, mqtt.Client) {
	log.Println("Inside Configure")

	//Prepare MQTT option
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", conData.HostName, conData.HostPort))
	opts.SetClientID(conData.ConnectionName)
	opts.SetUsername(conData.UserName)
	opts.SetPassword(conData.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	//Prepare connection
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return false, nil
	}
	return true, client
}

func Disconnect(client mqtt.Client) bool {
	log.Println("Inside Disconnect")
	client.Disconnect(250)
	return true
}

func Subscribe(client mqtt.Client, topic string, qos int) bool {
	log.Println("Inside Subscribe")
	client.Subscribe(topic, byte(qos), func(subclient mqtt.Client, msg mqtt.Message) {
		log.Printf("------ [%s] %s\n", msg.Topic(), string(msg.Payload()))
		//model.Messages <- msg.Topic() + "," + string(msg.Payload())

		payload := new(model.Payload)
		payload.Topic = msg.Topic()
		payload.Data = string(string(msg.Payload()))
		model.MessageChannel <- *payload

		// mqttData := &model.Payload{Topic: msg.Topic(), Data: string(msg.Payload())}
		// jsonData, err := json.Marshal(mqttData)
		// if err == nil {
		// 	log.Println("Json Data is...")
		// 	log.Println(jsonData)
		// 	payload := new(model.Payload)
		// 	payload.Topic = msg.Topic()
		// 	payload.Data = string(jsonData)
		// 	model.MessageChannel <- *payload
		// } else {
		// 	log.Printf("Error is: %s", err)
		// }
	})
	return true
}

func Unsubscribe(client mqtt.Client) bool {
	log.Println("Inside Unsubscribe")
	client.Unsubscribe()
	return true
}

func Publish(client mqtt.Client, topic string, qos byte, retained bool, payload interface{}) bool {
	log.Println("Inside Publish")
	client.Publish(topic, qos, retained, payload)
	return true
}

//Callback functions
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	options := client.OptionsReader()
	fmt.Println(options.ClientID() + " connected.")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	options := client.OptionsReader()
	fmt.Printf(options.ClientID()+" connection lost: %v\n", err)
}
