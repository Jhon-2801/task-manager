package routers

import (
	"encoding/json"
	"net/http"
	bd "v3/bd/task"
	"v3/models"
)

func UpDateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	Id := r.URL.Query().Get("id")

	if err != nil {
		http.Error(w, "Error al mandar los datos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(task.Name) == 0 {
		http.Error(w, "El campo nombre es requerido", http.StatusBadRequest)
		return
	}
	if len(task.Date) == 0 {
		http.Error(w, "El campo nombre es requerido", http.StatusBadRequest)
		return
	}
	if len(task.Description) == 0 {
		http.Error(w, "El campo nombre es requerido", http.StatusBadRequest)
		return
	}
	if len(task.Status) == 0 {
		http.Error(w, "El campo nombre es requerido", http.StatusBadRequest)
		return
	}
	if task.Id_user == 0 {
		http.Error(w, "El campo nombre es requerido", http.StatusBadRequest)
		return
	}

	bd.UpDateTaskBd(Id, task)
	w.WriteHeader(http.StatusOK)
}
