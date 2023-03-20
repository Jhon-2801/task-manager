package models

type Task struct {
	Id          int
	Name        string `json:"name"`
	Progress    int    `json:"progress"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Id_user     int    `json:"id_user"`
	Status      string `json:"status"`
}
