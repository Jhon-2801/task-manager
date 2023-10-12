package handlers

import (
	"net/http"

	"github.com/Jhon-2801/task-manager/core/service"
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
		RegisterUser Controller
		LoginUser    Controller
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

func MakeEnponints(s service.Service) EndPoints {
	return EndPoints{
		RegisterUser: makeRegisterUser(s),
	}
}

func makeRegisterUser(s service.Service) Controller {
	return func(c *gin.Context) {
		var req RegisterReq
		c.BindJSON(&req)
		if service.Service.IsValidMail(s, req.Mail) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "email is not valid "})
		}
		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "name is required"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "password is required"})
			return
		}
		c.IndentedJSON(http.StatusAccepted, req)
	}
}
