package models

type User struct {
	Id       int
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
