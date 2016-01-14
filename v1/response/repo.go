package response

type RepoInfo struct {
	ID           uint     `json:"id"`
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	FolderPath   string   `json:"folderPath"`
	Url          string   `json:"url"`
	IsRepo       bool     `json:"isRepo"`
	IsAuthorized bool     `json:"isAuthorized"`
	Description  string   `json:"description"`
	References   []string `json:"references"`
	Locked       bool     `json:"locked"`
}

type RepoListResponse struct {
	Repositories []RepoInfo `json:"repositories"`
}

type RepoCommit struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Date    int64  `json:"date"`
}

type RepoDiffFile struct {
	Type    string `json:"type"`
	OldId   string `json:"oldId"`
	OldName string `json:"oldName"`
	NewId   string `json:"newId"`
	NewName string `json:"newName"`
	Patch   string `json:"patch"`
}

type RepoCommitResponse struct {
	BasicResponse
	Commit RepoCommit     `json:"commit"`
	Files  []RepoDiffFile `json:"files"`
}

type RepoCommitsResponse struct {
	BasicResponse
	Commits []RepoCommit `json:"commits"`
}

type RepoInfoResponse struct {
	BasicResponse
	Info RepoInfo `json:"info"`
}

type RepoFile struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

type RepoFilesResponse struct {
	BasicResponse
	Files        []RepoFile `json:"files"`
	ParentTreeId string     `json:"parentTreeId"`
}
type RepoFileContentResponse struct {
	BasicResponse
	Size    int64  `json:"size"`
	Content string `json:"content"`
}
