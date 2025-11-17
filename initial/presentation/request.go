package presentation

import "github.com/RelaxContinent/go-echo-ddd-practice/initial/application"

type InitialRequest struct {
	UserID int `query:"user_id" validate:"required,numeric,min=1,max=10"`
}

// ToApplicationParams
// Presentation層のRequestをApplication層のParamsに変換する
func (r *InitialRequest) ToApplicationParams() *application.InitialParams {
	return &application.InitialParams{
		UserID: r.UserID,
	}
}
