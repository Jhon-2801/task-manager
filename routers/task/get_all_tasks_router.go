package routers

import (
	"encoding/json"
	"net/http"
	bd "v3/bd/task"
)

// Manda todas las tareas del usuario
func GetTasksRoutes(w http.ResponseWriter, r *http.Request) {
	UserID := r.URL.Query().Get("id_user")

	if len(UserID) == 0 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
	}

	tasksUser, err := bd.GetAllTaskUser(UserID)

	if err != nil {
		http.Error(w, "Error al intentar buscar las tareas del usuario"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tasksUser)
}
