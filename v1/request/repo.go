package request

type RepoListRequest struct {
	SubPath string `json:"subPath"`
}

type RepoCommitsRequest struct {
	RepoId    int    `json:"repoId"`
	Branch    string `json:"branch"`
	Start     int    `json:"start"`
	Count     int    `json:"count"`
	Ascending bool   `json:"ascending"`
}
