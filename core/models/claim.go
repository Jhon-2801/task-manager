package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	Mail string `json: "mail"`
	jwt.StandardClaims
}
