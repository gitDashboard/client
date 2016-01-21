package request

import (
	"github.com/gitDashboard/client/v1/admin/response"
)

type CreateRepoRequest struct {
	FolderID    uint   `json:"folderId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdatePermissionsRequest struct {
	Permissions []response.Permission `json:"permissions"`
}

type RepoMoveRequest struct {
	FolderID uint   `json:"folderId"`
	DestName string `json:"destName"`
}
