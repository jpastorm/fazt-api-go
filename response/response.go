package response

import "echi/models"

type ResponseUser struct {
	Users     []models.User `json:"data"`
	TotalUser int64         `json:"total_user"`
	Response  string        `json:"response"`
	Status    int64         `json:"status"`
	Pages     Page          `json:"pages"`
}
type Page struct {
	NextPage     int64 `json:"nextpage"`
	PreviousPage int64 `json:"previouspage"`
	TotalPage    int64 `json:"totalpage"`
}
