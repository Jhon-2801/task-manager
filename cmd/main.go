package main

import (
	"log"

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

	router := gin.Default()

	router.POST("/login", gin.HandlerFunc(userEnd.LoginUser))
	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))

	router.POST("/create", gin.HandlerFunc(taskEnd.CreateTask))

	router.Run(":8081")
}
