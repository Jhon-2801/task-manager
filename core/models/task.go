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
	Due_date    time.Time `json:"due_date"`
	UserID      int       `json:"id_user"`
	Status      string    `json:"status"`
	Create_at   time.Time `json:"create_at"`
	Update_at   time.Time `json:"update_at"`
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
