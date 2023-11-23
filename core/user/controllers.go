package user

import (
	"net/http"

	"github.com/Jhon-2801/task-manager/core/jwt"
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
		First_Name string `json:"first_name"`
		Last_Name  string `json:"last_name"`
		Mail       string `json:"mail"`
		Password   string `json:"password"`
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
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid "})
			return
		}
		_, err := s.GetUserByMail(req.Mail)
		if err == nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email already exists"})
			return
		}
		if req.First_Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "first name is required"})
			return
		}
		if req.Last_Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "last name is required"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "password is required"})
			return
		}

		err = s.Register(req.First_Name, req.Last_Name, req.Mail, req.Password)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}
		user, _ := s.GetUserByMail(req.Mail)
		jwtKey, err := jwt.GeneroJWT(user)
		req.Password = ""

		c.IndentedJSON(http.StatusCreated, gin.H{"status": 201, "user": req, "token": jwtKey})
	}
}

func makeLoginUser(s Service) Controller {
	return func(c *gin.Context) {
		var req LoginReq
		c.BindJSON(&req)
		if !Service.IsValidMail(s, req.Mail) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid "})
			return
		}
		user, err := s.GetUserByMail(req.Mail)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email no exists"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "password is required"})
			return
		}
		valid, err := s.ValidPassword(req.Mail, req.Password)
		if !valid {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid password"})
			return
		}
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}
		jwtKey, err := jwt.GeneroJWT(user)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}
		user.Password = ""
		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "data": user, "token": jwtKey})
	}
}

func makeGetAllUser(s Service) Controller {
	return func(c *gin.Context) {
		users, err := s.GetAllUser()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}

		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "users": users})
	}
}
