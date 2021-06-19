package user

import (
	"fmt"
	"time"
)

type UserLoginFormatter struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	DateBirth     string `json:"date_birth"`
	Address       string `json:"address"`
	Authorization string `json:"authorization"`
}

type UserFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	DateBirth string `json:"date_birth"`
	Address   string `json:"address"`
}

func ConvertDateBirth(dateBirth time.Time) string {
	var dateFormat = fmt.Sprintf("%v-%v-%v", dateBirth.Year(), int(dateBirth.Month()), dateBirth.Day())

	return dateFormat
}

func UserFormat(user User) UserFormatter {
	return UserFormatter{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		DateBirth: ConvertDateBirth(user.DateBirth),
		Address:   user.Address,
	}
}

func UserLoginFormat(user User, token string) UserLoginFormatter {
	return UserLoginFormatter{
		ID:            user.ID,
		Name:          user.Name,
		Email:         user.Email,
		DateBirth:     ConvertDateBirth(user.DateBirth),
		Address:       user.Address,
		Authorization: token,
	}
}

func UserFormats(users []User) []UserFormatter {
	var userFormats []UserFormatter

	for _, user := range users {
		userFormats = append(userFormats, UserFormat(user))
	}

	return userFormats
}
