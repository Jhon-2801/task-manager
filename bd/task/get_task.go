package bd

import (
	"log"
	"v3/models"
)

func GetTask(taskID int) models.Task {
	tasks, err := GetAllTask()

	if err != nil {
		log.Fatal("Error al consultar la base de datos")
	}

	var task models.Task

	for tasks.Next() {
		err = tasks.Scan(&task.Id, &task.Name, &task.Description, &task.UserID, &task.Date, &task.Status, &task.Progress)

		if err != nil {
			log.Fatal("error al consular la base de datos", err)
			return task
		}
		if taskID == task.Id {
			return task
		}
	}
	return task
}
