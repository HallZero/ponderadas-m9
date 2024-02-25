# Ponderada 4 - Integração simulador com HiveMQ

## 1. Objetivo

Criar conexão com autenticação para o simulador de dispositivos IoT utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho.

## 2. Enunciado

Nessa atividade, deve-se desenvolver a integração entre o simulador desenvolvido nas duas primeiras atividades ponderadas e um cluster configurado no HiveMQ. Para tal, deve-se garantir que o simulador é capaz de se comunicar utilizando autenticação em camada de transporte (TLS).


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
   go run .
   ```

### Subscriber:

1. Entre no diretório _subscriber_:
   ```bash
   cd subscriber
   ```
2. Rode o comando:
   ```bash
   go run .
   ```

### Modificações

Um novo método para criação de um cliente utilizando credenciais de autenticação num cluster em nívem do HiveMQ:

```go
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
```

Essas informações são buscadas diretamente de um arquivo .env (Não disponibilizado no repositório)

```go
err := godotenv.Load("./.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	newBroker := os.Getenv("BROKER_ADDR")
	user := os.Getenv("HIVE_USER")
	pswd := os.Getenv("HIVE_PSWD")
```

```go
client := DefaultClient.CreateClientWithAuth(newBroker, DefaultClient.IdPublisher, DefaultClient.Handler, user, pswd)
```

### Testes

Não foram modificados neste momento

## Vídeo de demonstração

[Simulador de dispositivos IoT](https://youtu.be/IQjbnmZvMmI)

Comparação entre o Tráfego de dados antes e após o vídeo:
<img src="../static/Screenshot from 2024-02-25 14-41-00.png" />
<img src="../static/Screenshot from 2024-02-25 14-39-17.png" />
