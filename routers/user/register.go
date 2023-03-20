package routers

import (
	"encoding/json"
	"net/http"
	bd "v3/bd/user"
	"v3/models"
	"v3/utilities"
)

// Registra usuario
func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos", http.StatusBadRequest)
		return
	}

	if len(user.Name) == 0 {
		http.Error(w, "El nombre del usuario es requerido", http.StatusBadRequest)
		return
	}
	if !utilities.IsValidMail(user.Mail) {
		http.Error(w, "El mail no es valido", http.StatusBadRequest)
		return
	}
	if len(user.Password) < 8 {
		http.Error(w, "El password tiene que tener mas de 8 caracteres", http.StatusBadRequest)
		return
	}

	itFoundUser, _ := bd.CheckExistUser(user.Mail)

	if itFoundUser {
		http.Error(w, "El usuario ya esta registrado", 400)
		return
	}

	bd.InsertRegister(user)
	w.WriteHeader(http.StatusCreated)

}
