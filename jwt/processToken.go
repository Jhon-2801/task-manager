package jwt

import (
	"errors"
	bd "v3/bd/user"
	"v3/models"

	tk "github.com/dgrijalva/jwt-go"
)

var Mail string

// Valida el JWT del usuario
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	claims := &models.Claim{}

	privateKey, _, err := LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")

	if err != nil {
		return claims, false, "", err
	}

	tkn, err := tk.ParseWithClaims(token, claims, func(t *tk.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err == nil {
		encontrado, _ := bd.CheckExistUser(claims.Mail)
		if encontrado {
			Mail = claims.Mail
		}
		return claims, encontrado, Mail, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token Invalido")
	}

	return claims, false, "", err
}
