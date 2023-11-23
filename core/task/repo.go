package task

import (
	"fmt"

	"github.com/Jhon-2801/task-manager/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(task *models.Task) error
		GetUserById(id string) error
		GetAllTask(id string) ([]models.Task, error)
	}
	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Create(task *models.Task) error {
	if err := repo.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUserById(id string) error {
	user := models.User{}
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetAllTask(id string) ([]models.Task, error) {
	var tasks []models.Task
	err := repo.db.Where("user_id = ?", id).Find(&tasks)

	fmt.Println(tasks)
	if err.Error != nil {
		return nil, err.Error
	}
	return tasks, nil
}
