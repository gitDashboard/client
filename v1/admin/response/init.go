package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type Permission struct {
	Users    []publicResponse.User `json:"users"`
	Types    []string              `json:"types"`
	Ref      string                `json:"ref"`
	Granted  bool                  `json:"granted"`
	Position uint                  `json:"position"`
}

type GetPermissionsResponse struct {
	publicResponse.BasicResponse
	Permissions []Permission `json:"permissions"`
}

type UpdatePermissionsResponse struct {
	publicResponse.BasicResponse
}
