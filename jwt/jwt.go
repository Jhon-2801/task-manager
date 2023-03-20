package jwt

import (
	"log"
	"time"
	"v3/models"

	"github.com/dgrijalva/jwt-go"
)

// Genera el JWT al usuario
func GeneraJwt(mail string) (string, error) {
	privateBytes, _, err := LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")

	if err != nil {
		log.Fatal("No se pudo cargar los certificados:", err)
	}
	claim := models.Claim{
		Mail: mail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "John",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(privateBytes)

	if err != nil {
		return "", err
	}
	return tokenStr, err
}
