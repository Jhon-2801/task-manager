package jwt

import (
	"v3/bd"

	"github.com/dgrijalva/jwt-go"
)

func GeneraJwt(mail string) (string, error) {
	miClave := []byte("jhon0128")

	user := bd.GetUser(mail)

	payload := jwt.MapClaims{
		"mail":    user.Mail,
		"name":    user.Name,
		"id_user": user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, err
}

func GetUser(mail string) {
	panic("unimplemented")
}
