package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       string
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
