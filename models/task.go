package models

type Task struct {
	Id          int
	Name        string `json:"name"`
	Description string `json:"description"`
}
