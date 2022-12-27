package presenter

import (
	"myapp/model"
)

type UserResponseWrapper struct {
	User *model.User `json:"user"`
}

type ListUserResponseWrapper struct {
	Users []model.User `json:"users"`
	Meta  interface{}  `json:"meta"`
}

type SignUpResponseWrapper struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
