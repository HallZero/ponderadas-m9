# Ponderada 1 - Simulador de dispositivos IoT

## 1. Objetivo

Criar um simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho.

## 2. Enunciado

A primeira atividade do módulo tem como objetivo a criação de um simulador de dispositivos IoT capaz de enviar informações em um tópico com o formato de dados consistente com os seguintes dispositivos de exemplo:

_Sensor de Radiação Solar_

_SPS30_

_MiCS-6814_

Escolha ao menos um desses sensores e estude o seu datasheed para reproduzir o que seria uma mensagem MQTT gerada a partir dos dados aferidos por eles. Foque na reprodução fidedigna de taxa de comunicação, unidade e ordem de grandeza das medições. Utilize alguma técnica/biblioteca para que suas mensagens simuladas não sejam todas iguais (mas todas dentro das especificações do componente).

Embora não haja o requerimento de criar testes automatizados, o simulador deve apresentar evidências objetivas de funcionamento.

## Lançando a aplicação

Para utilizar os módulos, clone o repositório, se ainda não o tiver feito e entre no diretório _ponderada1_:

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
   go run publisher.go
   ```

### Subscriber:

1. Entre no diretório _subscriber_:
   ```bash
   cd subscriber
   ```
2. Rode o comando:
   ```bash
   go run subscriber.go
   ```

## Publisher

Para permitir a criação de diversas instâncias no contexto de simulação, foi pensado em uma estrutura de dados que pudesse representar clientes MQTT como sensores que enviassem suas leituras como dados.

```go
type Sensor struct {
	Name        string
	Latitude    float64
	Longitude   float64
	Measurement float64
	Rate        int
	Unit        string
}
```

A partir dessa estrutura, pode-se criar instâncias que se conectam com um broker (nesse código, optou-se por utilizar o hivemq) e publicam suas leituras no tópico 'sensor/NomeDoSensor'. Para simular as leituras, gerou-se números aleatórios.

```go
for {
		for _, sensor := range sensors {

			topic := "sensors/" + sensor.Name

			sensor.Measurement = (rand.Float64() * (maxSensorRange - minSensorRange)) + minSensorRange

			payload, _ := sensor.ToJSON()

			token := client.Publish(topic, 0, false, payload)

			token.Wait()

			fmt.Printf("Published message: %s\n", payload)

			time.Sleep(time.Duration(sensor.Rate) * time.Second)

		}
	}
```

### Testes

Para assegurar que as instâncias dos sensores são criadas corretamente, utilizo uma função auxiliar que permite verificar a igualdade entre duas instâncias da estrutura. Se ambas forem iguais, o sensor foi criado corretamente.

```go
t.Run("Create new Sensor", func(t *testing.T) {
		sensor := NewSensor("Sensor1", 51.0, 0.0, 0.0, 60, "μg/m³")
		compare := &Sensor{Name: "Sensor1", Latitude: 51.0, Longitude: 0.0, Measurement: 0.0, Rate: 60, Unit: "μg/m³"}

		if !reflect.DeepEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})
```

## Subscriber

Da mesma forma, o subscriber conecta-se ao broker e subscreve-se em um tópico de interesse.

```go
if token := client.Subscribe("sensors/Sensor1", 1, nil); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    return
}

fmt.Println("Subscriber running...")
select {}
```

### Testes

Para assegurar que as instâncias dos sensores conseguem se conectar ao broker, utilizo uma função auxiliar que permite verificar a conectividade.

```go
t.Run("Subscription to topic", func(t *testing.T) {
		client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		if token := client.Subscribe("sensors/SPS30", 1, nil); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
			return
		}

		t.Log("Subscribed successfully to Topic")
	})
```
## Vídeo de demonstração

[Simulador de dispositivos IoT](https://youtu.be/J5b6GMtt3Q8)