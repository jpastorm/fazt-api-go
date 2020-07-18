package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"fazt-api-go/controllers"
)

func Connect() {
	//token := "QWERTY"
	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb+srv://user:Sistemas.2020@cluster0.eo4if.mongodb.net/faztdb?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	db := client.Database("faztdb")
	controllers.UserCollection(db)
	return
}
