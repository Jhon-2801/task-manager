package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	Id          string
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Dates       time.Time `json:"date"`
	UserID      int       `json:"id_user"`
	Status      bool      `json:"status"`
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
