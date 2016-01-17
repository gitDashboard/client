package request

import (
	"github.com/gitDashboard/client/v1/admin/response"
)

type CreateFolderRequest struct {
	Path string `json:"path"`
}

type CreateRepoRequest struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}

type UpdatePermissionsRequest struct {
	Permissions []response.RepoPermission `json:"permissions"`
}

type RepoMoveRequest struct {
	DestPath string `json:"destPath"`
	DestName string `json:"destName"`
}
