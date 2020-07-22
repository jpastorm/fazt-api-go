package controllers

import (
	"context"
	"echi/config"
	"echi/errorGo"
	"echi/models"
	"echi/response"
	"echi/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = config.Connect().Collection("user")

func AllUser(c echo.Context) error {

	channel := make(chan []byte)

	var response string
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")

	if &limit == nil || len(limit) == 0 || &page == nil || len(page) == 0 {

		go getAllUser(c, channel)
		response = string(<-channel)

	} else {

		if _, err := strconv.Atoi(limit); err == nil {
			fmt.Println("LOOKS LIKE A NUMBER")
			lp, _ := strconv.Atoi(limit)
			pp, _ := strconv.Atoi(page)
			l := int64(lp)
			p := int64(pp)
			go getLimitUser(c, l, p, channel)
		} else {
			lp := 10
			pp := 1
			l := int64(lp)
			p := int64(pp)
			go getLimitUser(c, l, p, channel)
		}

		response = string(<-channel)
	}
	return c.String(http.StatusOK, response)
}

func getAllUser(c echo.Context, channel chan []byte) {

	var users []models.User

	cursor, err := collection.Find(context.TODO(), bson.D{})

	errorGo.PanicError(err)

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)

		errorGo.LogFatalError(err)

		users = append(users, user)
	}

	resJson := &response.ResponseUser{
		Users:    users,
		Response: "Success",
		Status:   200,
	}

	response, _ := json.Marshal(resJson)
	channel <- response

}
func getLimitUser(c echo.Context, limit int64, page int64, channel chan []byte) {

	var previusPage int64
	if page == 0 || page < 0 || page == 1 {
		page = 1
		previusPage = 1
	} else {
		previusPage = page
	}

	limit, skips := utils.GetPages(limit, page)

	findOpts := options.Find()
	findOpts.SetLimit(limit)
	findOpts.SetSkip(skips)

	var users []models.User

	cursor, err := collection.Find(context.TODO(), bson.D{}, findOpts)
	count, err := collection.CountDocuments(context.TODO(), bson.M{})

	var totalPage int64
	totalPage = count / limit
	if totalPage%limit != 0 {
		totalPage++
	}

	errorGo.PanicError(err)

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)

		errorGo.LogFatalError(err)

		users = append(users, user)
	}
	pages := &response.Page{
		NextPage:     page + 1,
		PreviousPage: previusPage - 1,
		TotalPage:    totalPage,
	}
	resJson := &response.ResponseUser{
		Users:     users,
		TotalUser: count,
		Response:  "Success",
		Status:    200,
		Pages:     *pages,
	}
	response, _ := json.Marshal(resJson)
	channel <- response
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
