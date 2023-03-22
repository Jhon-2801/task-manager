package bd

import (
	"database/sql"
	"log"
	"v3/bd"
)

// Trae a todos los usuarios
func GetAllUsers() (*sql.Rows, error) {

	db, err := bd.GetConnectionBd()

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
