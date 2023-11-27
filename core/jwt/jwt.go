package jwt

import (
	"os"
	"time"

	"github.com/Jhon-2801/task-manager/core/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GeneroJWT(user models.User) (string, error) {
	_ = godotenv.Load()
	privateKey := os.Getenv("PRIVATE_KEY")
	privateKeyByte := []byte(privateKey)
	payload := jwt.MapClaims{
		"email":      user.Email,
		"first_name": user.First_Name,
		"last_name":  user.Last_Name,
		"id":         user.Id,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKeyByte)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
