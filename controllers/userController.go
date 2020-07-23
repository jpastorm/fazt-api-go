package controllers

import (
	"context"
	"echi/config"
	"echi/errorGo"
	"echi/models"
	"echi/response"
	"echi/utils"
	"encoding/json"
	"net/http"
	"strconv"

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
		if utils.IsInteger(limit) && utils.IsInteger(page) {
			//fmt.Println("LOOKS LIKE A NUMBER")
			l, _ := strconv.ParseInt(limit, 10, 64)
			p, _ := strconv.ParseInt(page, 10, 64)
			go getLimitUser(c, l, p, channel)
		} else {
			l := int64(10)
			p := int64(1)
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

	var ResponseUserId response.ResponseUserId

	if err := c.Bind(&user); err != nil {
		return c.JSON(500, &user)
	}
	var u models.IUser = user
	ok, id := u.Create()

	if ok {
		ResponseUserId = response.ResponseUserId{id, "El usuario fue creado", 201}
	} else {
		ResponseUserId = response.ResponseUserId{id, "No pudo crearse el usuario", 500}
	}
	return c.JSON(http.StatusCreated, ResponseUserId)

}

func LoginUser(c echo.Context) error {

	var user models.User

	var responseBody response.Response
	if err := c.Bind(&user); err != nil {
		return c.JSON(500, &user)
	}

	var u models.IUser = user
	res := u.Login()

	if res {
		responseBody = response.Response{"Logueado con exito", 200}
	} else {
		responseBody = response.Response{"No pudo loguearse", 500}
	}
	return c.JSON(http.StatusCreated, responseBody)

}

func SearchUser(c echo.Context) error {
	value := c.Param("value")
	var user models.User
	var ResponseManyUsers response.ResponseManyUsers
	var u models.IUser = user
	success, res := u.Search(value)
	if success {
		ResponseManyUsers = response.ResponseManyUsers{res, "Exito en busqueda", 200}
	} else {
		ResponseManyUsers = response.ResponseManyUsers{res, "Error de Busqueda", 500}
	}
	return c.JSON(http.StatusCreated, ResponseManyUsers)
}
