package routers

import (
	"net/http"
	bd "v3/bd/task"
	"v3/utilities"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	TaskId := r.URL.Query().Get("id_task")
	UserID := r.URL.Query().Get("id_user")

	if len(TaskId) == 0 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
	}
	if len(UserID) == 0 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
	}

	if !utilities.IsValidActionUser(TaskId, UserID) {
		http.Error(w, "Usuario invalido", http.StatusBadRequest)
		return
	}

	bd.DeleteTaskBD(TaskId)
	w.WriteHeader(http.StatusOK)

}
