package helper

import (
	model "rmzstartup/model/entity"

	"github.com/google/uuid"
)

type UserFormatter struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
}

func FormatUser(user model.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}
