package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type CreateFolderResponse struct {
	publicResponse.BasicResponse
}

type SetFolderAdminsResponse struct {
	publicResponse.BasicResponse
}
