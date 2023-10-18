package task

import (
	"github.com/Jhon-2801/task-manager/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(task *models.Task) error
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
