package service

import (
	"regexp"

	Repository "github.com/Jhon-2801/task-manager/core/respositories"
)

type (
	Service interface {
		Register(name, mail, password string) error
		IsValidMail(mail string) bool
	}
	service struct {
		repo Repository.Repository
	}
)

func NewService(repo Repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Register(name, mail, password string) error {
	return nil
}

func (s service) IsValidMail(mail string) bool {
	validMail := regexp.MustCompile("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$")
	return validMail.MatchString(mail)
}
