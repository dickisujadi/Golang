package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mongodb/mongo-go-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	// set client option
	clientOptions := options.Client().ApplyURI("your cluster endpoint")

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_rest_api").Collection("books")

	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json: "status"`
	ErrorMessage string `json: "message"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())

	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}