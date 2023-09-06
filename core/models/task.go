package models

type Task struct {
	Id          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	UserID      int    `json:"id_user"`
	Status      string `json:"status"`
}
