package routers

import (
	"encoding/json"
	"net/http"
	bd "v3/bd/task"
)

// Manda todas las tareas del usuario
func GetTasksRoutes(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")

	if len(Id) == 0 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	tasksUser, err := bd.GetAllTaskUser(Id)

	if err != nil {
		http.Error(w, "Error al intentar buscar las tareas del usuario"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tasksUser)
}
