package respositories

import (
	"github.com/Jhon-2801/task-manager/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Register(user *models.User) error
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

// Register implements Repository.
func (*repo) Register(user *models.User) error {

	return nil
}
