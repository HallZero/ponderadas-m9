package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sensor struct {
	Name        string  `bson:"name"`
	Latitude    float64 `bson:"latitude"`
	Longitude   float64 `bson:"longitude"`
	Measurement float64 `bson:"measurement"`
	Rate        int     `bson:"rate"`
	Unit        string  `bson:"unit"`
}


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
