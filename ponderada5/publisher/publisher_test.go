package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	DefaultClient "mqtt/src/common"
)

var client = DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

func TestMain(t *testing.T) {
	t.Run("Create new Sensor", func(t *testing.T) {
		sensor := NewSensor("Sensor1", 51.0, 0.0, 0.0, 60, "μg/m³")
		compare := &Sensor{Name: "Sensor1", Latitude: 51.0, Longitude: 0.0, Measurement: 0.0, Rate: 60, Unit: "μg/m³"}

		if !reflect.DeepEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})

	t.Run("Generating JSON file to payload", func(t *testing.T) {
		sensor := NewSensor("SPS30", 51.0, 0.0, 0.0, 1, "μg/m³")

		got, err := sensor.ToJSON()

		var transformed map[string]interface{}

		json.Unmarshal([]byte(got), &transformed)

		if err != nil {
			t.Errorf("Error generating JSON: %v", err)
		}

		want := map[string]interface{}{
			"Name":        "SPS30",
			"Latitude":    51.0,
			"Longitude":   0.0,
			"Measurement": 0.0,
			"Rate":        1,
			"Unit":        "μg/m³",
		}

		// May change this later. Map comparison is quite confusing (reflect.DeepEqual() returns false)
		if !(fmt.Sprint(transformed) == fmt.Sprint(want)) {
			t.Errorf("Unexpected JSON output.\nGot: %v\nWant: %v", transformed, want)
		}

	})

	t.Run("Test QoS - eg if the message was published by the broker", func(t *testing.T) {

		payload := "Hello, Broker!"

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		token := client.Publish("sensors", 1, false, payload)

		if token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		t.Log("Broker received message with QoS 1!")
	})

	t.Run("Publish Rate check", func(t *testing.T) {

		const size = 5
		const tolerance = 10000
		const frequence = 1000000 // 1 milisecond
		const especulated_time = frequence + tolerance

		start_time := time.Now().Nanosecond()

		for i := 0; i < size; i++ {

			// payload := rand.Float32()

			token := client.Publish("sensors", 1, false, "Testing sensor Rate Publish")

			if token.Wait() && token.Error() != nil {
				t.Error(token.Error())
			}

		}

		end_time := time.Now().Nanosecond()

		mean_sensor_time := (end_time - start_time) / size

		if mean_sensor_time > especulated_time {
			t.Fatalf("Time is bigger than especulated. Wanted: %d, but got: %d", especulated_time, mean_sensor_time)
		}

		t.Log("Time is within the especulated range")

	})
}
