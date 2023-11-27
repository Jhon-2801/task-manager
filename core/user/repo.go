package user

import (
	"github.com/Jhon-2801/task-manager/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Register(user *models.User) error
		GetAllUser() ([]models.User, error)
		GetUserByMail(email string) (models.User, error)
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

func (repo *repo) Register(user *models.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUserByMail(email string) (models.User, error) {
	user := models.User{}
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *repo) GetAllUser() ([]models.User, error) {
	var user []models.User

	tx := repo.db.Select("id", "first_name", "last_name", "email").Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
