package model

//go:generate go run github.com/mailru/easyjson/easyjson

//easyjson:json
type AcceptInviteRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}
