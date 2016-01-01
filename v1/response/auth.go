package response

type AuthorizationResponse struct {
	Authorized bool `json:"authorized"`
}

type LoginResponse struct {
	BasicResponse
	JWT string `json:"jwt_token"`
}
