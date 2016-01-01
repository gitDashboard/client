package response

type RepoInfo struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	IsRepo       bool   `json:"isRepo"`
	IsAuthorized bool   `json:"isAuthorized"`
	Description  string `json:"description"`
}

type RepoListResponse struct {
	Repositories []RepoInfo `json:"repositories"`
}

type RepoCommit struct {
	Author  string `json:"author"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Date    int64  `json:"date"`
}

type RepoCommitsResponse struct {
	BasicResponse
	Commits []RepoCommit `json:"commits"`
}
