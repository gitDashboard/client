package request

type FindEventRequest struct {
	RepoID      uint     `json:"repoId"`
	User        string   `json:"user"`
	Reference   string   `json:"reference"`
	Description string   `json:"description"`
	Levels      []string `json:"level"`
	Since       int64    `json:"since"`
	To          int64    `json:"to"`
	Type        string   `json:"type"`
	First       int     `json:"first"`
	Count       int     `json:"count"`
}
