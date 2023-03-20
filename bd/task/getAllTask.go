package bd

import (
	"database/sql"
	"log"
	"v3/bd"
)

// Trae todas las tareas
func GetAllTask() (*sql.Rows, error) {

	db, err := bd.GetConnectionBd()

	if err != nil {
		log.Fatal("Error al establecer una conexión a la base de datos", err)
	}
	defer db.Close()

	tasks, err := db.DB().Query("SELECT id_task, name_task, description_task, user_id_user, dat_task, state_task, progreso FROM task")

	if err != nil {
		log.Fatal("error al consular la base de datos", err)
	}

	return tasks, err

}
