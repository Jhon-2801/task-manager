package user

import (
	"crypto/rsa"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/Jhon-2801/task-manager/core/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type (
	Service interface {
		Register(first_name, last_Name, mail, password string) error
		IsValidMail(mail string) bool
		GetUserByMail(mail string) (models.User, error)
		EncryptPassword(password string) (string, error)
		ValidPassword(mail, password string) (bool, error)
		GenerateJWT(mail string) string
	}
	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Register(first_name, last_Name, mail, password string) error {

	password, err := s.EncryptPassword(password)

	if err != nil {
		return err
	}
	user := models.User{
		First_Name: first_name,
		Last_Name:  last_Name,
		Mail:       mail,
		Password:   password,
	}

	err = s.repo.Register(&user)

	if err != nil {
		return err
	}

	return nil
}

func (s service) IsValidMail(mail string) bool {
	validMail := regexp.MustCompile("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$")
	return validMail.MatchString(mail)
}

func (s service) EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s service) GetUserByMail(mail string) (models.User, error) {
	user, err := s.repo.GetUserByMail(mail)
	return user, err
}

func (s service) ValidPassword(mail, password string) (bool, error) {
	user, err := s.repo.GetUserByMail(mail)

	if err != nil {
		return true, err
	}
	passwordByte := []byte(password)
	passwordDB := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(passwordDB, passwordByte)

	if err != nil {
		return false, err
	}

	return true, nil
}

func init() {
	privateKeyBytes, err := os.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("Could not read private file")
	}

	publicKeyBytes, err := os.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("Could not read public file")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatal("Could not parse privateKey")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
}

func (s service) GenerateJWT(mail string) string {
	claim := models.Claim{
		Mail: mail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Could not sign token")
	}
	return result
}
