package request

type CreateFolderRequest struct {
	ParentID    uint   `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
