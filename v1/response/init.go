package response

type Error struct {
	Code    string `json:"errorCode"`
	Message string `json:"message"`
}

func init() {
	initErrors()
}

var NoRepositoryFoundError, PermissionDeniedError, FatalError, NoUserFoundError, AuthenticationFailedError Error

func initErrors() {
	FatalError = Error{Code: "Fatal Error", Message: "FatalError:"}
	PermissionDeniedError = Error{Code: "PermissionDenied", Message: "Permission Denied"}
	NoRepositoryFoundError = Error{Code: "NoRepoFound", Message: "No Repository Found"}
	NoUserFoundError = Error{Code: "NoUserFound", Message: "No User Found"}
	AuthenticationFailedError = Error{Code: "AuthenticationFailed", Message: "Authentication failed"}
}

type BasicResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}
