package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	//guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	//"time"
)


type User struct {
	ID string `json:"_id`
	Nickname string `json:"nickName`
	Firstname string `json:"firstName`
	Lastname string `json:"lastName`
	Email string `json:"email`
	Password string `json:"password`
}

// DATABASE INSTANCE
var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("user")
}
func GetAllUsers(c *gin.Context) {
	todos := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var todo User
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})
	return
}

