package models

import (
	"context"
	"echi/config"
	"echi/errorGo"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUser interface {
	Login() bool
	Create() (bool, string)
	Search(value string) (bool, []User)
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname  string             `json:"nickName,omitempty"`
	Firstname string             `json:"firstName,omitempty"`
	Lastname  string             `json:"lastName,omitempty"`
	Email     string             `json:"email,omitempty"`
	Password  string             `json:"password,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
}

var collection = config.Connect().Collection("user")

func (u User) Login() bool {
	var result User
	filter := bson.M{"nickname": u.Nickname, "password": u.Password}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	fmt.Println(result)
	return true
}
func (u User) Create() (bool, string) {

	result, err := collection.InsertOne(context.TODO(), u) //InsertOne solo nos devuelve un error o el id
	if err != nil {
		return false, "null"
	}
	id := result.InsertedID.(primitive.ObjectID).Hex() //Conversion a string

	return true, id
}

func (u User) Search(value string) (bool, []User) {
	var users []User
	//filter := bson.M{"nickname": value}
	filter := bson.D{{"nickname", primitive.Regex{Pattern: value, Options: ""}}}
	result, _ := collection.Find(context.TODO(), filter)

	for result.Next(context.TODO()) {
		err := result.Decode(&u)
		errorGo.LogFatalError(err)
		users = append(users, u)
	}
	if len(users) == 0 {
		return false, users
	}

	return true, users
}
