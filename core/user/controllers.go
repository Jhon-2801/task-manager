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
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	UserRes struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	RegisterReq struct {
		FirstName string `form:"first_name"`
		LastName  string `form:"last_name"`
		Email     string `form:"email"`
		Password  string `form:"password"`
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
		err := c.ShouldBind(&req)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 500, "message": err})
			return
		}
		if req.Email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is required"})
			return
		}
		if !Service.IsValidMail(s, req.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid "})
			return
		}
		_, err = s.GetUserByMail(req.Email)
		if err == nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email already exists"})
			return
		}
		if req.FirstName == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "first name is required"})
			return
		}
		if req.LastName == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "last name is required"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the password must be greater than 7 characters"})
			return
		}

		err = s.Register(req.FirstName, req.LastName, req.Email, req.Password)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}
		user, _ := s.GetUserByMail(req.Email)
		jwtKey, err := jwt.GeneroJWT(user)

		data := UserRes{
			Id:        user.Id,
			FirstName: user.First_Name,
			LastName:  user.Last_Name,
			Email:     user.Email,
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"status": 201, "user": data, "token": jwtKey})
	}
}

func makeLoginUser(s Service) Controller {
	return func(c *gin.Context) {
		var req LoginReq
		err := c.ShouldBind(&req)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 500, "message": err})
			return
		}
		if req.Email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is required"})
			return
		}
		if !Service.IsValidMail(s, req.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid "})
			return
		}
		user, err := s.GetUserByMail(req.Email)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email no exists"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the password must be greater than 7 characters"})
			return
		}
		valid, err := s.ValidPassword(req.Email, req.Password)
		if !valid {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "invalid password"})
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

		data := UserRes{
			Id:        user.Id,
			FirstName: user.First_Name,
			LastName:  user.Last_Name,
			Email:     user.Email,
		}
		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "data": data, "token": jwtKey})
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
