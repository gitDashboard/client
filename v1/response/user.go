package response

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type UsersResponse struct {
	BasicResponse
	Users []User `json:"users"`
}
