package jwt

import (
	"os"
	"strings"

	"github.com/Jhon-2801/task-manager/core/models"

	"github.com/dgrijalva/jwt-go"
)

func ProcessToken(tk string) (bool, error) {
	privateKey := os.Getenv("PRIVATE_KEY")

	privateKeyByte := []byte(privateKey)
	var claims models.Claim

	slpitToken := strings.Split(tk, "Bearer")
	if len(slpitToken) != 2 {
		return false, nil
	}

	tk = strings.TrimSpace(slpitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(t *jwt.Token) (interface{}, error) {
		return privateKeyByte, nil
	})

	if err == nil {
		return true, nil
	}
	if !tkn.Valid {
		return false, nil
	}
	return false, err
}
