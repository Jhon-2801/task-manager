package jwt

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Jhon-2801/task-manager/core/models"
	"github.com/joho/godotenv"

	"github.com/dgrijalva/jwt-go"
)

func ProcessToken(tk string) (bool, error) {
	_ = godotenv.Load()
	privateKey := os.Getenv("PRIVATE_KEY")

	privateKeyByte := []byte(privateKey)
	var claims models.Claim

	slpitToken := strings.Split(tk, "Bearer")
	if len(slpitToken) != 2 {
		return false, errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(slpitToken[1])

	fmt.Println(tk)
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(1)
		return privateKeyByte, nil
	})

	if err == nil {
		fmt.Println(2)
		return true, nil
	}
	fmt.Println(1)
	if !tkn.Valid {
		return false, errors.New("token Inválido")
	}
	return false, err
}
