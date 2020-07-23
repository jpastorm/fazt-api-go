package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	//uri := "mongodb+srv://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + "?retryWrites=true&w=majority"
	urilocal := "mongodb://localhost:27017"
	// Database Config
	clientOptions := options.Client().ApplyURI(urilocal) //AQUI COLOCAMOS LA URL DE CONEXION
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
	db := client.Database(dbName)
	//controllers.UserCollection(db)
	return db
}
