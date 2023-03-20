package bd

import (
	"log"
	"strconv"
	"v3/bd"
	"v3/models"
)

func UpDateTaskBd(id string, task models.Task) {
	bd, err := bd.GetConnectionBd()

	if err != nil {
		log.Fatal("Error al intentar conectarse con la base de datos" + err.Error())
		return
	}

	defer bd.Close()

	idTask, _ := strconv.ParseInt(id, 10, 0)

	sentencia, err := bd.CommonDB().Prepare("UPDATE task SET name_task = ?, description_task = ?, user_id_user = ?, dat_task = ?, state_task = ?, progreso = ? WHERE id_task = ?")

	if err != nil {
		log.Fatal("Error al actualizar datos" + err.Error())
		return
	}

	defer sentencia.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentencia.Exec(task.Name, task.Description, task.Id_user, task.Date, task.Status, task.Progress, idTask)
	if err != nil {
		log.Fatal("Error al actualizar datos" + err.Error())
		return
	}
}
