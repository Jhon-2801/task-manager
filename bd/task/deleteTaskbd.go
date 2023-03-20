package bd

import (
	"log"
	"strconv"
	"v3/bd"
)

func DeleteTaskBD(Id string) error {
	bd, err := bd.GetConnectionBd()
	if err != nil {
		log.Fatal("Error al establecer una conexión a la base de datos", err)
		return err
	}
	defer bd.Close()

	idTask, _ := strconv.ParseInt(Id, 10, 0)

	sentencia, err := bd.CommonDB().Prepare("DELETE FROM task WHERE id_task = ?")
	if err != nil {
		return err
	}
	defer sentencia.Close()

	_, err = sentencia.Exec(idTask)
	if err != nil {
		return err
	}
	return nil
}
