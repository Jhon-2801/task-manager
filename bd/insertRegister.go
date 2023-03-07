package bd

import (
	"log"
	"v3/models"
)

func InsertRegister(user models.User) {
	db, _ := GetConnectionBd()

	defer db.Close()

	user.Password = EncriptarPassword(user.Password)

	insertarUser, err := db.DB().Prepare("INSERT INTO users (name_user, mail, password) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal("Error al insertar usuario", err)
	}
	defer insertarUser.Close()
	// Ejecutar sentencia, un valor por cada '?'
	insertarUser.Exec(user.Name, user.Mail, user.Password)
}
