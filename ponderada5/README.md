# Ponderada 5 - Integração do simulador com Metabase

## 1. Objetivo

Integração entre o simulador desenvolvido nas últimas atividades com Metabase

## 2. Enunciado

Nessa atividade, deve-se desenvolver a integração entre o simulador desenvolvido nas três primeiras atividades e um dashboard desenvolvido usando o Metabase, com persistência de dados em um banco de dados a sua escolha.


## Lançando a aplicação

Para utilizar os módulos, clone o repositório, se ainda não o tiver feito e entre no diretório _ponderada5_:

```bash
    git clone https://github.com/HallZero/ponderadas-m9.git
    cd ponderadas-m9/ponderada5
```

### Metabase e MongoDB

Nessa aplicação, utilizamos tanto o Metabase quanto o MongoDB em containers. Para lançar os dois conjuntamente, vá para a o diretório _ponderada5_ e rode o comando:

```bash
docker compose up -d
```

> Nota: Para fazer a configuração da database no Metabase, no campo Host, coloque o nome do container "mongodb", a porta 27017 e Database name "teste_banco".

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

### API:

1. Entre no diretório _api_:
   ```bash
   cd api
   ```
2. Rode o comando:
   ```bash
   go run .
   ```

### Modificações

Criação de uma API REST simples para a injeção de dados no Banco de Dados.

```go
func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "")
}

func postData(c *gin.Context) {

	var newSensor Sensor

	if err := c.BindJSON(&newSensor); err != nil {
		return
	}

	col := connectToDatabase()
	insertIntoDatabase(col, newSensor)
}

func connectToDatabase() *mongo.Collection {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("teste_banco")

	collection := db.Collection("test_collection")

	return collection

}

func insertIntoDatabase(collection *mongo.Collection, data Sensor) {

	insertionResult, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Documento inserido com o ID: %v\n", insertionResult.InsertedID)

}

func main() {
	router := gin.Default()
	router.GET("/data", getData)
	router.POST("/data", postData)

	router.Run("localhost:8080")

}
```

### Testes

Não foram modificados neste momento

## Vídeo de demonstração

[ntegração Metabase + Simulação utilizando MongoDB](https://youtu.be/FC6o3KCbTmA)
