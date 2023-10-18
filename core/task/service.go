package task

import (
	"fmt"
	"time"

	"github.com/Jhon-2801/task-manager/core/models"
)

type (
	Service interface {
		Create(name, descrip, date string) error
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

func (s service) Create(name, descrip, date string) error {
	newDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}
	task := &models.Task{
		Name:        name,
		Description: descrip,
		Dates:       newDate,
	}
	err = s.repo.Create(task)
	if err != nil {
		fmt.Println("El error esta aqui", err)
		return err
	}
	return nil
}
