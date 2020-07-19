// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID        string `json:"id" bson:"_id"`
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
