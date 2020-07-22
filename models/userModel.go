package models

import (
	"context"
	"echi/config"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUser interface {
	Login() bool
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname  string             `json:"nickName"`
	Firstname string             `json:"firstName"`
	Lastname  string             `json:"lastName"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `json:"created_at"`
}

var collection = config.Connect().Collection("user")

func (u User) Login() bool {

	var result User
	/*
		fmt.Println("DESDE EL MODELO: " + u.Nickname)
		fmt.Println("DESDE EL MODELO:" + u.Password) */

	filter := bson.M{"nickname": u.Nickname, "password": u.Password}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	fmt.Println(result)

	return true
}
