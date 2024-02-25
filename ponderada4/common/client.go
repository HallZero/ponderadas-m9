package common

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const Broker = "broker.hivemq.com:1883"
const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"

var Handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
	return
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func CreateClient(broker string, id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)

	return mqtt.NewClient(opts)
}

func CreateClientWithAuth(broker string, id string, callback_handler mqtt.MessageHandler, user string, password string) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, 8883))
	opts.SetClientID(id)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(callback_handler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return mqtt.NewClient(opts)
}
