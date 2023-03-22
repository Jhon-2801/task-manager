package routers

import (
	"encoding/json"
	"net/http"
	bd "v3/bd/task"
	"v3/models"
	"v3/utilities"
)

func UpDateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	TaskId := r.URL.Query().Get("id_task")
	UserID := r.URL.Query().Get("id_user")

	if err != nil {
		http.Error(w, "Error al mandar los datos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(TaskId) == 0 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
	}
	if len(UserID) == 0 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
	}

	if len(task.Name) == 0 {
		http.Error(w, "El campo nombre es requerido", http.StatusBadRequest)
		return
	}
	if len(task.Date) == 0 {
		http.Error(w, "El campo date es requerido", http.StatusBadRequest)
		return
	}
	if len(task.Description) == 0 {
		http.Error(w, "El campo description es requerido", http.StatusBadRequest)
		return
	}
	if len(task.Status) == 0 {
		http.Error(w, "El campo status es requerido", http.StatusBadRequest)
		return
	}

	if !utilities.IsValidActionUser(TaskId, UserID) {
		http.Error(w, "Usuario invalido", http.StatusBadRequest)
		return
	}
	bd.UpDateTaskBd(TaskId, task)
	w.WriteHeader(http.StatusOK)
}
