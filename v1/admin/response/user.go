package response

import (
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type UserUpdateResponse struct {
	publicResponse.BasicResponse
}

type UserDeleteResponse struct {
	publicResponse.BasicResponse
}
