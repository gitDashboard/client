package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type Group struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Users       []User `json:"users"`
}

type GroupsResponse struct {
	publicResponse.BasicResponse
	Groups []Group `json:"groups"`
}

type GroupUpdateResponse struct {
	publicResponse.BasicResponse
}

type GroupDeleteResponse struct {
	publicResponse.BasicResponse
}
