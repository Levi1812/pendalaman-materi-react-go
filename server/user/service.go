package user

import (
	"backendGolang/auth"
	"errors"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input UserRegisterInput) (UserFormatter, error)
	LoginUser(input UserLoginInput) (UserLoginFormatter, error)
	GetAllUser() ([]UserFormatter, error)
	GetUserByID(userID string) (UserFormatter, error)
}

type service struct {
	repository  Repository
	authService auth.Service
}

func NewService(repository Repository, authService auth.Service) *service {
	return &service{repository, authService}
}

func (s *service) RegisterUser(input UserRegisterInput) (UserFormatter, error) {
	genPass, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	dateSplit := strings.Split(input.DateBirth, "-")

	year, _ := strconv.Atoi(dateSplit[0])
	month, _ := strconv.Atoi(dateSplit[1])
	day, _ := strconv.Atoi(dateSplit[2])

	dateTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	var newUser = User{
		Name:      input.Name,
		Address:   input.Address,
		DateBirth: dateTime,
		Email:     input.Email,
		Password:  string(genPass),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.repository.Create(newUser)
	if err != nil {
		return UserFormatter{}, err
	}

	formatter := UserFormat(user)

	return formatter, nil

}
func (s *service) LoginUser(input UserLoginInput) (UserLoginFormatter, error) {

	checkUser, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return UserLoginFormatter{}, err
	}

	if checkUser.ID == 0 || len(checkUser.Name) <= 1 {
		return UserLoginFormatter{}, errors.New("user email / password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(input.Password))

	if err != nil {
		return UserLoginFormatter{}, errors.New("user email / password invalid")
	}

	token, _ := s.authService.GenerateToken(checkUser.ID)

	formatter := UserLoginFormat(checkUser, token)

	return formatter, nil
}

func (s *service) GetAllUser() ([]UserFormatter, error) {
	users, err := s.repository.GetAll()

	if err != nil {
		return []UserFormatter{}, err
	}

	formatters := UserFormats(users)

	return formatters, nil
}

func (s *service) GetUserByID(userID string) (UserFormatter, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return UserFormatter{}, err
	}

	formatter := UserFormat(user)

	return formatter, nil
}
