package presentation

import (
	"strconv"

	"github.com/RelaxContinent/go-echo-ddd-practice/initial/application"
)

type InitialResponse struct {
	Message string `json:"message"`
}

// NewInitialResponse
// Application層のResultをPresentation層のResponseに変換する
func NewInitialResponse(result *application.InitialResult) *InitialResponse {
	message :=
		"User Info: " +
			"ID=" + strconv.Itoa(int(result.User.ID)) +
			", Name=" + result.User.Name +
			", UserName=" + result.User.UserName +
			", Email=" + result.User.Email +
			", Website=" + result.User.Website
	return &InitialResponse{
		Message: message,
	}
}
