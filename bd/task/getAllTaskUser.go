package bd

import (
	"log"
	"strconv"
	"v3/models"
)

// Trae todas las tareas del usuario
func GetAllTaskUser(Id string) ([]models.Task, error) {
	var task models.Task

	tasks, err := GetAllTask()

	if err != nil {
		log.Fatal("Error al traer las tareas", err)
	}
	tasksUser := []models.Task{}

	idUser, _ := strconv.ParseInt(Id, 10, 0)

	for tasks.Next() {
		err = tasks.Scan(&task.Id, &task.Name, &task.Description, &task.Id_user, &task.Date, &task.Status, &task.Progress)
		if err != nil {
			log.Fatal("error al consular la base de datos", err)
			return tasksUser, err
		}
		if int(idUser) == task.Id_user {
			tasksUser = append(tasksUser, task)
		}
	}

	return tasksUser, err
}
