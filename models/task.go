package models

import (
	"time"
)

type Task struct {
	Id          int
	Name        string    `json:"name"`
	Progress    int       `json:"progress"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Id_user     int       `json:"id_user"`
}
