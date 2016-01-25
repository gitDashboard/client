package request

import (
	basicresponse "github.com/gitDashboard/client/v1/response"
)

type CreateFolderRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SetFolderAdminsRequest struct {
	Admins []basicresponse.User `json:"admins"`
}
