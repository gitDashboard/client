package request

import (
	"github.com/gitDashboard/client/v1/misc"
)

type RepoListRequest struct {
	SubPath string `json:"subPath"`
}

type RepoCommitsRequest struct {
	RepoId int    `json:"repoId"`
	Branch string `json:"branch"`
	Start  int    `json:"start"`
	Count  int    `json:"count"`
}

type RepoFilesRequest struct {
	RepoId  int    `json:"repoId"`
	RefName string `json:"refName"`
	Parent  string `json:"parent"`
}

type RepoEventRequest struct {
	RepositoryPath string `json:"path"`
	Type           string `json:"type"`
	Level          misc.EventLevel
	User           string `json:"user"`
	Description    string `json:"description"`
	Reference      string `json:"reference"`
}
