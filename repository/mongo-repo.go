package repository

import (


	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	model2 "github.com/jpastorm/gqlgen-todos/graph/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Save(user *model2.User)
	FindAll() []*model2.User
}
type database struct {
	client *mongo.Client
}


const 	DATABASE = "faztdb"
const 	COLLECTION = "user"


func New() *database {

	MONGODB := "mongodb+srv://user:Sistemas.2020@cluster0.eo4if.mongodb.net/faztdb?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(MONGODB)

	clientOptions = clientOptions.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")

	return &database{
		client: dbClient,
	}
}
func (db *database) Save(user *model2.User) {
	//collection := db.client.Database["faztdb"].Collection("user")
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_,err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}
func (db *database) FindAll() []*model2.User {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor ,err := collection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model2.User
	for cursor.Next(context.TODO()){
		var u *model2.User
		err := cursor.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, u)
	}
	return result
}
