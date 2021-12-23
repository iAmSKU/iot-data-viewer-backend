package model

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttConfiguration struct {
	ClientName string
	HostName   string
	HostPort   string
	UserName   string
	Password   string
	Topic      string
}

type Connectivity interface {
	Configure() (bool, *mqtt.ClientOptions)
	Connect() bool
	Disconnect() bool
	Subscribe(topic string, qos byte) bool
	Unsubscribe() bool
	Publish(topic string, qos byte, retained bool, payload interface{}) bool
}

func (conData MqttConfiguration) Configure() (bool, *mqtt.ClientOptions) {
	log.Println("Inside Configure")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", conData.HostName, conData.HostPort))
	opts.SetClientID(conData.ClientName)
	opts.SetUsername(conData.UserName)
	opts.SetPassword(conData.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return true, opts
}

func (conData MqttConfiguration) Connect(opts *mqtt.ClientOptions) (bool, mqtt.Client) {
	log.Println("Inside Connect")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return false, client
	}
	return true, client
}

func (conData MqttConfiguration) Disconnect(client mqtt.Client) bool {
	log.Println("Inside Disconnect")
	client.Disconnect(250)
	return true
}

func (conData MqttConfiguration) Subscribe(client mqtt.Client, topic string, qos byte) bool {
	log.Println("Inside Subscribe")
	client.Subscribe(topic, qos, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("------ [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
	return true
}

func (conData MqttConfiguration) Unsubscribe(client mqtt.Client) bool {
	log.Println("Inside Unsubscribe")
	client.Unsubscribe()
	return true
}

func (conData MqttConfiguration) Publish(client mqtt.Client, topic string, qos byte, retained bool, payload interface{}) bool {
	log.Println("Inside Publish")
	client.Publish(topic, qos, retained, payload)
	return true
}

//Callback functions
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
