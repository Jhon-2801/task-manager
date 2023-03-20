package bd

import (
	"v3/models"

	"golang.org/x/crypto/bcrypt"
)

// Hace la validacion de la contraseña del usuario
func ValidatePassword(user models.User) bool {
	itFoundUser, password := CheckExistUser(user.Mail)

	if !itFoundUser {
		return false
	}
	passwordByte := []byte(user.Password)
	passwordDB := []byte(password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordByte)

	if err != nil {
		return false
	}

	return true
}
