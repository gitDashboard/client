package response

type FolderInfo struct {
	ID          uint   `json:"id"`
	ParentID    uint   `json:"parentId"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

type FolderListResponse struct {
	BasicResponse
	Folders []FolderInfo `json:"folders"`
}

type FolderGetResponse struct {
	BasicResponse
	Folder FolderInfo `json:"folder"`
}
