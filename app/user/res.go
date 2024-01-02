package user

import (
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/models"
)

type ResUser struct {
	ID       int		`json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Photos   []models.Photo
}

type ResUserToken struct {
	ResUser
	Token    string    `json:"token"`
}

func FormatUserResponse(user models.User, token string) interface{} {
	var formatter interface{}

	if token == "" {
		formatter = ResUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
	} else {
		userResponse := ResUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		formatter = ResUserToken{
			ResUser: userResponse,
			Token:        token,
		}
	}

	return formatter
}