package handlers

import (
	"net/http"
	"v3/middlew"
	"v3/routers"

	"github.com/gorilla/mux"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.ChequeoDb(routers.RegisterUser)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDb(routers.LoginUser)).Methods("POST")

	http.ListenAndServe(":3000", router)

}
