package response

import (
	"github.com/gitDashboard/client/v1/misc"
	publicResponse "github.com/gitDashboard/client/v1/response"
)

type Event struct {
	ID          uint            `json:"id"`
	RepoID      uint            `json:"repoId"`
	Type        string          `json:"type"`
	Level       misc.EventLevel `json:"level"`
	Started     int64           `json:"started"`
	Finished    int64           `json:"finished"`
	User        string          `json:"user"`
	Reference   string          `json:"reference"`
	Description string          `json:"description"`
}

type FindEventResponse struct {
	publicResponse.BasicResponse
	Events []Event `json:"events"`
}
