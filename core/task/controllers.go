package task

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
		CreateTask Controller
		GetAllTask Controller
	}
	CreateReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Due_date    string `json:"due_date"`
		UserID      string `json:"user_id"`
	}
)

func MakeEnponints(s Service) EndPoints {
	return EndPoints{
		CreateTask: makeCreateTask(s),
		GetAllTask: makeGetAllTask(s),
	}
}

func makeCreateTask(s Service) Controller {
	return func(c *gin.Context) {
		var req CreateReq
		c.BindJSON(&req)

		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "name is required"})
			return
		}
		if req.Due_date == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "due_date is required"})
			return
		}
		if req.UserID == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}

		err := s.GetUserById(req.UserID)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user not found", "err": err})
			return
		}

		newDate, err := time.Parse("2006-01-02", req.Due_date)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}

		if newDate.Before(time.Now()) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the date provided is before the current date"})
			return
		}

		err = s.Create(req.Name, req.Description, req.UserID, newDate)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"status": 201, "data": req})
	}
}

func makeGetAllTask(s Service) Controller {
	return func(c *gin.Context) {
		userId := c.Param("id")
		if userId == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}
		tasks, err := s.GetAllTask(userId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"status": 200, "data": tasks})

	}
}
