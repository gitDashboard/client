package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Type     string `json:"type"`
}

type UsersResponse struct {
	publicResponse.BasicResponse
	Users []User `json:"users"`
}
