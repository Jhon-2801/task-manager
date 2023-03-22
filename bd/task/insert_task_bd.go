package bd

import (
	"log"
	"time"
	"v3/bd"
	"v3/models"
)

// Inserta la tarea a la BD
func InsertTaskBd(task models.Task) {
	db, _ := bd.GetConnectionBd()

	defer db.Close()

	date, err := time.Parse("2006-01-02", task.Date)

	if err != nil {
		log.Fatal("Error", err)
		return
	}

	insertarUser, err := db.DB().Prepare("INSERT INTO task (name_task, description_task, user_id_user, dat_task, state_task, progreso) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error al insertar Tarea", err)
		return
	}
	defer insertarUser.Close()
	// Ejecutar sentencia, un valor por cada '?'
	insertarUser.Exec(task.Name, task.Description, task.UserID, date, task.Status, task.Progress)
}
