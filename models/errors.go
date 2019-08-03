package models

type errorMessage struct {
	Message string `json:"message"`
}

func (e errorMessage) Error() string {
	return e.Message
}

var UserNotFound = errorMessage{"User not found."}
var InvalidJson = errorMessage{"Invalid JSON."}
var InvalidCredentials = errorMessage{"Invalid credentials."}
var UserAlreadyExist = errorMessage{"User already exist."}
var InvalidToken = errorMessage{"Invalid token."}
var TokenExpired = errorMessage{"Token has expired"}
