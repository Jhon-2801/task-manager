package bd

import (
	"log"
)

func CheckExistUser(mail string) bool {

	var existMail string

	db, err := GetConnetionBd()

	if err != nil {
		log.Fatal("Fallo al conectarse con la base de datos", err)
	}
	defer db.Close()

	mails, err := db.DB().Query("SELECT mail FROM users")

	if err != nil {
		log.Fatal("error al consular la base de datos", err)
	}

	defer mails.Close()

	for mails.Next() {
		err = mails.Scan(&existMail)

		if err != nil {
			log.Fatal("error al consular la base de datos", err)
			return true
		}
		if mail == existMail {
			return true
		}
	}
	return false
}
