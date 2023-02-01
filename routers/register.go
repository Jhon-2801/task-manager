package routers

import (
	"encoding/json"
	"net/http"
	"v3/bd"
	"v3/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error en los datos recibidos", 400)
		return
	}

	if len(user.Name) == 0 {
		http.Error(w, "El nombre del usuario es requerido", 400)
		return
	}
	if len(user.Mail) == 0 {
		http.Error(w, "El mail del usuario es requerido", 400)

		return
	}
	if len(user.Password) < 8 {
		http.Error(w, "El password del usuario es requerido", 400)
		return
	}

	itFoundUser := bd.CheckExistUser(user.Mail)

	if itFoundUser {
		http.Error(w, "El usuario ya esta registrado", 400)
		return
	}

	bd.InsertRegister(user)

}
