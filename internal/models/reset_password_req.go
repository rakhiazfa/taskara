package models

type ResetPasswordReq struct {
	Token                string `json:"token" validate:"required,max=100"`
	Password             string `json:"password" validate:"required,min=8,max=100"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,min=8,max=100,eqfield=Password"`
}
