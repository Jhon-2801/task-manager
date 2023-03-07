package bd

import (
	"database/sql"
	"log"
)

func GetAllUsers() (*sql.Rows, error) {

	db, err := GetConnectionBd()

	if err != nil {
		log.Fatal("Error al establecer una conexión a la base de datos", err)
	}
	defer db.Close()

	users, err := db.DB().Query("SELECT mail, password, name_user, id_user FROM users")

	if err != nil {
		log.Fatal("error al consular la base de datos", err)
	}

	return users, err

}
