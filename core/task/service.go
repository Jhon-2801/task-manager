package task

import (
	"time"

	"github.com/Jhon-2801/task-manager/core/models"
)

type (
	Service interface {
		Create(name, descrip, userID string, dueDate time.Time) error
		GetUserById(id string) error
		GetAllTask(id string) ([]models.Task, error)
		UpDateTask(id, name, descrip, userID string, dueDate time.Time, status bool, create time.Time) (string, error)
		GetTaskById(id string) error
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
		Update_at:   time.Now(),
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
func (s service) GetTaskById(id string) error {
	err := s.repo.GetTaskById(id)
	if err != nil {
		return err
	}
	return nil
}
func (s service) GetAllTask(id string) ([]models.Task, error) {
	tasks, err := s.repo.GetAllTask(id)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s service) UpDateTask(id, name, descrip, userID string, dueDate time.Time, status bool, create time.Time) (string, error) {

	task := &models.Task{
		Id:          id,
		Name:        name,
		Description: descrip,
		Due_date:    dueDate,
		UserID:      userID,
		Status:      false,
		Create_at:   create,
		Update_at:   time.Now(),
	}

	err := s.repo.UpDateTask(task)
	if err != nil {
		return "", err
	}
	update := task.Update_at.Format("2006-01-02")
	return update, err
}
