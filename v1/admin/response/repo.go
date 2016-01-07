package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type CreateFolderResponse struct {
	publicResponse.BasicResponse
}

type CreateRepoResponse struct {
	publicResponse.BasicResponse
}

type RepoPermission struct {
	UserID    int64  `json:"userId"`
	UserName  string `json:"userName"`
	GroupID   int64  `json:"groupId"`
	GroupName string `json:"groupName"`
	Types     string `json:"types"`
	Ref       string `json:"ref"`
	Granted   bool   `json:"granted"`
}

type GetPermissionsResponse struct {
	publicResponse.BasicResponse
	Permissions []RepoPermission `json:"permissions"`
}

type UpdatePermissionsResponse struct {
	publicResponse.BasicResponse
}
