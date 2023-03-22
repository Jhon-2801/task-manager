package bd

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Encripta la contraseña
func EncriptarPassword(pass string) string {
	cost := 6

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		log.Fatal("Error al encriptar password")
	}
	pass = string(bytes)
	return pass
}
