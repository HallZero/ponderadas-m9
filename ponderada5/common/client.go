package common

import (
	"bytes"
	"fmt"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const Broker = "broker.hivemq.com:1883"
const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"

var Handler mqtt.MessageHandler = func(_ mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())

	bodyReader := bytes.NewReader(msg.Payload())

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/data", bodyReader)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // Create an HTTP client
    client := &http.Client{}

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    // Check the response status
    if resp.StatusCode != http.StatusOK {
        fmt.Println("Unexpected response status:", resp.StatusCode)
        return
    }

    // Handle the response if needed

    fmt.Println("Request sent successfully")
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
