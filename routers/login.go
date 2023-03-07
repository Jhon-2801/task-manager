package routers

import (
	"encoding/json"
	"net/http"
	"v3/bd"
	"v3/jwt"
	"v3/models"
	"v3/utilities"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if !utilities.IsValidMail(user.Mail) {
		http.Error(w, "El email no es valido", 400)
		return
	}

	if len(user.Password) < 8 {
		http.Error(w, "Contraseña no valida", 400)
		return
	}

	isValidUser := bd.ValidatePassword(user)

	if !isValidUser {
		http.Error(w, "El usuario no es valido", 400)
		return
	}

	token, err := jwt.GeneraJwt(user.Mail)

	response := models.ResponseToken{
		Token: token,
	}
	if err != nil {
		http.Error(w, "Error al generar JWT"+err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	w.Write([]byte("Login accepted!"))
}
