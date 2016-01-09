package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type UsersResponse struct {
	publicResponse.BasicResponse
	Users []User `json:"users"`
}

type UserUpdateResponse struct {
	publicResponse.BasicResponse
}

type UserDeleteResponse struct {
	publicResponse.BasicResponse
}
