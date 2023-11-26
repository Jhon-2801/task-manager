package main

import (
	"log"

	"github.com/Jhon-2801/task-manager/core/middleware"
	"github.com/Jhon-2801/task-manager/core/task"
	"github.com/Jhon-2801/task-manager/core/user"
	"github.com/Jhon-2801/task-manager/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectionBD()

	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := user.NewRepo(db)
	userSrv := user.NewService(userRepo)
	userEnd := user.MakeEnponints(userSrv)

	taskRepo := task.NewRepo(db)
	taskSrv := task.NewService(taskRepo)
	taskEnd := task.MakeEnponints(taskSrv)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/login", gin.HandlerFunc(userEnd.LoginUser))
	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))
	router.GET("/users", middleware.ValidToken, gin.HandlerFunc(userEnd.GetAllUser))

	router.POST("/create", middleware.ValidToken, gin.HandlerFunc(taskEnd.CreateTask))
	router.GET("/tasks/:id", middleware.ValidToken, gin.HandlerFunc(taskEnd.GetAllTask))
	router.PATCH("/update/:id", middleware.ValidToken, gin.HandlerFunc(taskEnd.UpDateTask))

	router.Run(":8080")
}
