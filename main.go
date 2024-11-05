// main.go
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vishalc412/go-mongo-crud/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	// Initialize MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://admin:password123@localhost:27017/?authSource=admin")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	router := mux.NewRouter()
	userController := controllers.NewUserController(client)

	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/user", userController.CreateUser).Methods("POST")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
