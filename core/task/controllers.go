package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
		CreateTask Controller
	}
	CreateReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Date        string `json:"date"`
	}
)

func MakeEnponints(s Service) EndPoints {
	return EndPoints{
		CreateTask: makeCreateTask(s),
	}
}

func makeCreateTask(s Service) Controller {
	return func(c *gin.Context) {
		var req CreateReq
		c.BindJSON(&req)

		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "name is required"})
			return
		}
		if req.Date == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "date is required"})
			return
		}

		err := s.Create(req.Name, req.Description, req.Date)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusCreated, req)
	}
}
