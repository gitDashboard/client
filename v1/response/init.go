package response

type Error struct {
	Code    string `json:"errorCode"`
	Message string `json:"message"`
}

func init() {
	initErrors()
}

var NoRepositoryFoundError, PermissionDeniedError, FatalError, NoUserFoundError, NoGroupFoundError, AuthenticationFailedError, AlreadyExistError Error

func initErrors() {
	FatalError = Error{Code: "FatalError", Message: "FatalError:"}
	PermissionDeniedError = Error{Code: "PermissionDenied", Message: "Permission Denied"}
	NoRepositoryFoundError = Error{Code: "NoRepoFound", Message: "No Repository Found"}
	NoUserFoundError = Error{Code: "NoUserFound", Message: "No User Found"}
	NoGroupFoundError = Error{Code: "NoGroupFound", Message: "No Group Found"}
	AuthenticationFailedError = Error{Code: "AuthenticationFailed", Message: "Authentication failed"}
	AlreadyExistError = Error{Code: "AlreadyExist", Message: "Already exist"}
}

type BasicResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

type IBasicResponse interface {
	SetSuccess(success bool)
	SetError(respError Error)
}

func (resp *BasicResponse) SetSuccess(success bool) {
	resp.Success = success
}

func (resp *BasicResponse) SetError(respError Error) {
	resp.Error = respError
}
