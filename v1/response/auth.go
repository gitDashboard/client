package response

type AuthorizationResponse struct {
	Authorized bool `json:"authorized"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	JWT     string `json:"jwt_token"`
}
