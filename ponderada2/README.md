# Ponderada 2 - Teste de um simulador de dispositivos IoT

## 1. Objetivo

Criar testes para o simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho.

## 2. Enunciado

Utilizando o simulador de dispositivos IoT desenvolvido na atividade passada e utilizando os conceitos de TDD vistos no decorrer da semana, implemente testes automatizados para validar o simulador. Seus testes obrigatoriamente devem abordar os seguintes aspectos:

    Recebimento - garante que os dados enviados pelo simulador são recebidos pelo broker.
    Validação dos dados - garante que os dados enviados pelo simulador chegam sem alterações.
    Confirmação da taxa de disparo - garante que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.


## Lançando a aplicação

Para utilizar os módulos, clone o repositório, se ainda não o tiver feito e entre no diretório _ponderada2_:

```bash
    git clone https://github.com/HallZero/ponderadas-m9.git
    cd ponderadas-m9/ponderada1
```

### Publisher:

1. Entre no diretório _publisher_:

   ```bash
   cd publisher
   ```

2. Rode o comando:

   ```bash
   go test -v
   ```

### Subscriber:

1. Entre no diretório _subscriber_:
   ```bash
   cd subscriber
   ```
2. Rode o comando:
   ```bash
   go test -v
   ```

### Testes

Testes foram adicionados para avaliar o recebimento das mensagens pelo Broker usando o mecanismo de QoS, além da confirmação da taxa de disparo dos sensores

```go
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
```

Assegurando que as mensagens enviadas pelo publisher foram recebidas corretamente pelo subscriber:

```go
t.Run("Check Payload Integrity", func(t *testing.T) {
		
		publisher := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

		if token := publisher.Connect(); token.Wait() && token.Error() != nil {
			t.Fatal(token.Error())
		}

		defer publisher.Disconnect(250)

	// Initialize MQTT client for subscribing
		subscriber := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)
		
		if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
			t.Fatal(token.Error())
		}

		defer subscriber.Disconnect(250)

		// Subscribe to the topic
		topic := "sensors/SPS30"

		received := make(chan []byte)

		subscriber.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
			received <- message.Payload()
		})

		// Publish a message
		message := "test payload"
		publisher.Publish(topic, 1, false, message)

		// Wait for a short duration to receive the message
		select {
		case payload := <-received:
			if string(payload) != message {
				t.Errorf("Received payload %s, expected %s", payload, message)
			}
		case <-time.After(2 * time.Second):
			t.Error("Timeout: Did not receive the payload")
		}
	})
```
## Vídeo de demonstração

[Simulador de dispositivos IoT](https://youtu.be/wjfrcYlQbT8)