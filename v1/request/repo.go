package request

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
