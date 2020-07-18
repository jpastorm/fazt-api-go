package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)


type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nickname 	string `json:"nickName`
	Firstname 	string `json:"firstName`
	Lastname	string `json:"lastName`
	Email 		string `json:"email`
	Password 	string `json:"password`
	CreatedAt 	time.Time `json:"created_at"`
}
//TOKEN
var tokenUser = "QWERTY"
// DATABASE INSTANCE
var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("user")
}
func GetAllUsers(c *gin.Context) {

	token := c.Query("token")

	if  &token == nil || len(token) < 0 || token != tokenUser{
		log.Printf("No existe el token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "No tienes permisos",
		})
		return
	}
	users := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    users,
	})
	return
}
func CreateUser(c *gin.Context) {


	var user User
	c.BindJSON(&user)
	nickname := user.Nickname
	firstname := user.Firstname
	lastname := user.Lastname
	email := user.Email
	password := user.Password

	newUser := User{
		Nickname  :     nickname,
		Firstname :     firstname,
		Lastname  :     lastname,
		Email     :     email,
		Password  :     password,
		CreatedAt : 	time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "user created Successfully",
	})
	return
}
