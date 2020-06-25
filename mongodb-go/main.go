package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter,request *http.Request){
	response.Header().Add("content-type","application/json")
	ver person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyplotdeveloper").Collection("people")
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
}

func main() {
	fmt.Println("starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	http.ListenAndServe(":12345", router)

}
