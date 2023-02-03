package bd

import (
	"log"
)

func CheckExistUser(mail string) (bool, string) {

	var existMail string
	var password string

	db, err := GetConnetionBd()

	if err != nil {
		log.Fatal("Fallo al conectarse con la base de datos", err)
	}
	defer db.Close()

	mails, err := db.DB().Query("SELECT mail, password FROM users")

	if err != nil {
		log.Fatal("error al consular la base de datos", err)
	}

	defer mails.Close()

	for mails.Next() {
		err = mails.Scan(&existMail, &password)

		if err != nil {
			log.Fatal("error al consular la base de datos", err)
			return true, ""
		}
		if mail == existMail {

			return true, password
		}
	}
	return false, ""
}
