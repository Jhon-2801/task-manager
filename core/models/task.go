package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	Id          string
	Name        string `json:"name"`
	Description string `json:"description"`
	Due_date    string `json:"due_date"`
	UserID      int    `json:"id_user"`
	Status      string `json:"status"`
	Create_at   string `json:"create_at"`
	Update_at   string `json:"update_at"`
}

func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {

	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	return
}
