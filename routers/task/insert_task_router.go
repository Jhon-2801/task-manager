package routers

import (
	"encoding/json"
	"net/http"
	bdTask "v3/bd/task"
	"v3/models"
)

// Validad los datos de la nueva tarea e inserta en la base de datos
func InsertTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error al insetar los datos"+err.Error(), 400)
		return
	}

	if len(task.Name) == 0 {
		http.Error(w, "El campo es requerido", http.StatusBadRequest)
		return
	}

	if len(task.Date) == 0 {
		http.Error(w, "El campo es requerido", http.StatusBadRequest)
		return
	}

	task.Progress = 0
	task.Status = "New"

	bdTask.InsertTaskBd(task)

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
