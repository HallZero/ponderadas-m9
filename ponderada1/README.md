
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
	name        string
	latitude    float64
	longitude   float64
	measurement float64
	rate        int
	mqtt.Client
}
```

A partir dessa estrutura, pode-se criar instâncias que se conectam com um broker (nesse código, optou-se por utilizar o hivemq) e publicam suas leituras no tópico 'sensor/NomeDoSensor'. Para simular as leituras, gerou-se números aleatórios.

```go
topic := "sensors/" + sensor.name

for {
    sensor.measurement = rand.Float64()*5
    payload := strconv.FormatFloat(sensor.measurement, 'f', 2, 64)
    token := sensor.Publish(topic, 0, false, payload)
    token.Wait()
    fmt.Printf("Published message: %s\n", payload)
    time.Sleep(time.Duration(sensor.rate) * time.Second)
}
```

### Testes

Para assegurar que as instâncias dos sensores são criadas corretamente, utilizo uma função auxiliar que permite verificar a igualdade entre duas instâncias da estrutura. Se ambas forem iguais, o sensor foi criado corretamente.

```go
func TestNewSensor(t *testing.T) {
	t.Run("Create new Sensor", func(t *testing.T) {
		sensor := NewSensor("Sensor1", 51.0, 0.0, 0.0, 60)
		compare := Sensor{name: "Sensor1", latitude: 51.0, longitude: 0.0, measurement: 0.0, rate: 60}

		if structFieldsEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})
}
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

🚧 WIP 🚧