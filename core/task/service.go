package task

import (
	"time"

	"github.com/Jhon-2801/task-manager/core/models"
)

type (
	Service interface {
		Create(name, descrip, userID string, dueDate time.Time) error
		GetUserById(id string) error
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

func (s service) Create(name, descrip, userID string, dueDate time.Time) error {

	task := &models.Task{
		Id:          "",
		Name:        name,
		Description: descrip,
		Due_date:    dueDate,
		UserID:      userID,
		Status:      false,
		Create_at:   time.Now(),
		Update_at:   time.Time{},
	}
	err := s.repo.Create(task)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetUserById(id string) error {
	err := s.repo.GetUserById(id)
	if err != nil {
		return err
	}
	return nil
}
