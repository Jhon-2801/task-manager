package routers

import (
	"net/http"
	bd "v3/bd/task"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")

	if len(Id) == 0 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
	}

	bd.DeleteTaskBD(Id)
	w.WriteHeader(http.StatusOK)

}
