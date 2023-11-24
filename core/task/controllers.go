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
		UpDateTask Controller
	}
	CreateReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		UserID      string `json:"user_id"`
	}
	UpDateReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		UserID      string `json:"id_user"`
		Status      bool   `json:"status"`
		Create      string `json:"create_at"`
		UpDate      string
	}
)

func MakeEnponints(s Service) EndPoints {
	return EndPoints{
		CreateTask: makeCreateTask(s),
		GetAllTask: makeGetAllTask(s),
		UpDateTask: makeUpdateTask(s),
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
		if req.DueDate == "" {
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

		newDate, err := time.Parse("2006-01-02", req.DueDate)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
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

func makeUpdateTask(s Service) Controller {
	return func(c *gin.Context) {
		taskId := c.Param("id")
		if taskId == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "task_id is required"})
			return
		}
		err := s.GetTaskById(taskId)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user not found", "err": err})
			return
		}

		var req UpDateReq
		c.BindJSON(&req)

		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "name is required"})
			return
		}
		if req.Description == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "description is required"})
			return
		}
		if req.DueDate == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "due_date is required"})
			return
		}
		if req.Create == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "create_at is required"})
			return
		}
		if req.UserID == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}
		err = s.GetUserById(req.UserID)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user not found", "err": err})
			return
		}
		dueDate, err := time.Parse("2006-01-02", req.DueDate)
		create, err := time.Parse("2006-01-02", req.Create)

		update, err := s.UpDateTask(taskId, req.Name, req.Description, req.UserID, dueDate, req.Status, create)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		req.UpDate = update
		c.IndentedJSON(http.StatusOK, gin.H{"status": 200, "data": req})
	}
}
