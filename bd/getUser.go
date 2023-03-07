package bd

import (
	"log"
	"v3/models"
)

func GetUser(mail string) models.User {

	users, err := GetAllUsers()

	user := models.User{}

	if err != nil {
		return user
	}
	for users.Next() {
		err = users.Scan(&user.Mail, &user.Password, &user.Name, &user.Id)

		if err != nil {
			log.Fatal("error al consular la base de datos", err)
			return user
		}
		if mail == user.Mail {

			return user
		}
	}
	return user
}
