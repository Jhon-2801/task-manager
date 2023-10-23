package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
		RegisterUser Controller
		LoginUser    Controller
		GetAllUser   Controller
	}
	LoginReq struct {
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}
	RegisterReq struct {
		Name     string `json:"name"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	}
)

func MakeEnponints(s Service) EndPoints {
	return EndPoints{
		RegisterUser: makeRegisterUser(s),
		LoginUser:    makeLoginUser(s),
		GetAllUser:   makeGetAllUser(s),
	}
}

func makeRegisterUser(s Service) Controller {
	return func(c *gin.Context) {
		var req RegisterReq
		c.BindJSON(&req)
		if !Service.IsValidMail(s, req.Mail) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email is not valid "})
			return
		}
		_, err := s.GetUserByMail(req.Mail)
		if err == nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "the email already exists"})
			return
		}
		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "name is required"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "password is required"})
			return
		}

		err = s.Register(req.Name, req.Mail, req.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		token := s.GenerateJWT(req.Mail)

		c.IndentedJSON(http.StatusCreated, gin.H{"user": req, "token": token})
	}
}

func makeLoginUser(s Service) Controller {
	return func(c *gin.Context) {
		var req LoginReq
		c.BindJSON(&req)
		if !Service.IsValidMail(s, req.Mail) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email is not valid "})
			return
		}
		user, err := s.GetUserByMail(req.Mail)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "the email no exists"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "password is required"})
			return
		}
		valid, err := s.ValidPassword(req.Mail, req.Password)
		if !valid {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
			return
		}
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		token := s.GenerateJWT(req.Mail)

		c.IndentedJSON(http.StatusAccepted, gin.H{"user": user, "token": token})
	}
}

func makeGetAllUser(s Service) Controller {
	return func(c *gin.Context) {
		users, err := s.GetAllUser()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusAccepted, gin.H{"users": users})
	}
}
