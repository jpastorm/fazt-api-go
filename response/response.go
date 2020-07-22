package response

import "echi/models"

type ResponseUser struct {
	Users     []models.User `json:"data"`
	TotalUser int64         `json:"total_user,omitempty" bson:"total_user,omitempty"`
	Response  string        `json:"response"`
	Status    int64         `json:"status"`
	Pages     Page          `json:"pages,omitempty" bson:"pages,omitempty"`
}
type Page struct {
	NextPage     int64 `json:"nextpage,omitempty" bson:"nextpage,omitempty"`
	PreviousPage int64 `json:"previouspage,omitempty" bson:"previouspage,omitempty"`
	TotalPage    int64 `json:"totalpage,omitempty" bson:"totalpage,omitempty"`
}
type Response struct {
	Message string `json:"Message,omitempty" bson:"Message,omitempty"`
	Status  int    `json:"Status,omitempty" bson:"Status,omitempty"`
}