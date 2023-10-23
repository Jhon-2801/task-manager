package user

import (
	"github.com/Jhon-2801/task-manager/core/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Register(user *models.User) error
		GetAllUser() (*gorm.DB, error)
		GetUserByMail(mail string) (models.User, error)
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

func (repo *repo) GetUserByMail(mail string) (models.User, error) {
	user := models.User{}
	err := repo.db.Where("mail = ?", mail).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *repo) GetAllUser() (*gorm.DB, error) {
	user := models.User{}

	users := repo.db.Find(&user)
	if users.Error != nil {
		return nil, users.Error
	}
	return users, nil
}
