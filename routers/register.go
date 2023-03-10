package routers

import (
	"encoding/json"
	"net/http"
	"v3/bd"
	"v3/models"
	"v3/utilities"
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
	if !utilities.IsValidMail(user.Mail) {
		http.Error(w, "El mail no es valido", 400)
		return
	}
	if len(user.Password) < 8 {
		http.Error(w, "El password tiene que tener mas de 8 caracteres", 400)
		return
	}

	itFoundUser, _ := bd.CheckExistUser(user.Mail)

	if itFoundUser == true {
		http.Error(w, "El usuario ya esta registrado", 400)
		return
	}

	bd.InsertRegister(user)

}
