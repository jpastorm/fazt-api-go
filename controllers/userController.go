package controllers

import (
	"context"
	"echi/config"
	"echi/errorGo"
	"echi/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type content struct {
	Users    []models.User `json:"data"`
	Response string        `json:"response"`
	Status   int           `json:"Status"`
}

var collection = config.Connect().Collection("user")

func AllUser(c echo.Context) error {

	chanel := make(chan []byte)

	var response string

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	if &limit == nil || len(limit) == 0 || &page == nil || len(page) == 0 {

		go getAllUser(c, chanel)
		response = string(<-chanel)

	} else {

		var l int64
		var p int64
		fmt.Sscan(limit, &l)
		fmt.Sscan(page, &p)
		go getLimitUser(c, l, p, chanel)
		response = string(<-chanel)
	}
	return c.String(http.StatusOK, response)
}

func getAllUser(c echo.Context, chanel chan []byte) {

	var users []models.User

	cursor, err := collection.Find(context.TODO(), bson.D{})

	errorGo.PanicError(err)

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)

		errorGo.LogFatalError(err)

		users = append(users, user)
	}

	resJson := &content{
		Users:    users,
		Response: "Success",
		Status:   200,
	}
	response, _ := json.Marshal(resJson)
	chanel <- response

}
func getLimitUser(c echo.Context, limit int64, page int64, chanel chan []byte) {
	skips := limit * (page - 1)
	var users []models.User
	findOpts := options.Find()
	findOpts.SetLimit(limit)
	findOpts.SetSkip(skips)

	cursor, err := collection.Find(context.TODO(), bson.D{}, findOpts)

	errorGo.PanicError(err)

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)

		errorGo.LogFatalError(err)

		users = append(users, user)
	}

	resJson := &content{
		Users:    users,
		Response: "Success",
		Status:   200,
	}
	response, _ := json.Marshal(resJson)
	chanel <- response
}
func CreateUser(c echo.Context) error {

	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(500, &user)
	}
	nickname := user.Nickname
	firstname := user.Firstname
	lastname := user.Lastname
	email := user.Email
	password := user.Password

	newUser := models.User{
		Nickname:  nickname,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	collection.InsertOne(context.TODO(), newUser)
	return c.JSON(http.StatusCreated, newUser)
}
