package handlers

import (
	"net/http"
	"v3/middlew"
	routersTask "v3/routers/task"
	routersUser "v3/routers/user"

	"github.com/gorilla/mux"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.ChequeoDb(routersUser.RegisterUser)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDb(routersUser.LoginUser)).Methods("POST")
	router.HandleFunc("/tasks", middlew.ChequeoDb(middlew.ValidToken(routersTask.GetTasksRoutes))).Methods("GET")
	router.HandleFunc("/tasks", middlew.ChequeoDb(middlew.ValidToken(routersTask.InsertTask))).Methods("POST")
	router.HandleFunc("/task_delete", middlew.ChequeoDb(middlew.ValidToken(routersTask.DeleteTask))).Methods("DELETE")
	router.HandleFunc("/update_task", middlew.ChequeoDb(middlew.ValidToken(routersTask.UpDateTask))).Methods("PUT")

	http.ListenAndServe(":3000", router)

}
