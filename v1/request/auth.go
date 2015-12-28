package request

type AuthorizationRequest struct {
	Username       string `json:"username"`
	RepositoryPath string `json:"path"`
	RefName        string `json:"refName"`
	Operation      string `json:"operation"`
}
