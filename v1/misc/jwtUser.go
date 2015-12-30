package misc

type JWTUser struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
}
