package response

type ShortUserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type FolderInfo struct {
	ID             uint            `json:"id"`
	ParentID       uint            `json:"parentId"`
	Name           string          `json:"name"`
	Path           string          `json:"path"`
	Description    string          `json:"description"`
	Admins         []ShortUserInfo `json:"admins"`
	ExtendedAdmins []ShortUserInfo `json:"extAdmins"`
}

type FolderListResponse struct {
	BasicResponse
	Folders []FolderInfo `json:"folders"`
}

type FolderGetResponse struct {
	BasicResponse
	Folder FolderInfo `json:"folder"`
}
