package user

import (
	"regexp"

	"github.com/Jhon-2801/task-manager/core/models"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service interface {
		Register(first_name, last_Name, email, password string) error
		IsValidMail(email string) bool
		GetAllUser() ([]models.User, error)
		GetUserByMail(email string) (models.User, error)
		EncryptPassword(password string) (string, error)
		ValidPassword(email, password string) (bool, error)
	}
	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Register(first_name, last_Name, email, password string) error {

	password, err := s.EncryptPassword(password)

	if err != nil {
		return err
	}
	user := models.User{
		First_Name: first_name,
		Last_Name:  last_Name,
		Email:      email,
		Password:   password,
	}

	err = s.repo.Register(&user)

	if err != nil {
		return err
	}

	return nil
}

func (s service) IsValidMail(email string) bool {
	validMail := regexp.MustCompile("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$")
	return validMail.MatchString(email)
}

func (s service) EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s service) GetUserByMail(email string) (models.User, error) {
	user, err := s.repo.GetUserByMail(email)
	return user, err
}

func (s service) ValidPassword(email, password string) (bool, error) {
	user, err := s.repo.GetUserByMail(email)

	if err != nil {
		return true, err
	}
	passwordByte := []byte(password)
	passwordDB := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(passwordDB, passwordByte)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s service) GetAllUser() ([]models.User, error) {
	users, err := s.repo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, err
}
