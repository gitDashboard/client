package response

type AuthorizationResponse struct {
	Authorized bool `json:"authorized"`
	Locked     bool `json:"locked"`
}

type LoginResponse struct {
	BasicResponse
	JWT string `json:"jwt_token"`
}
