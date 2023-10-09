package main

import (
	"log"

	con "github.com/Jhon-2801/task-manager/core/controllers"
	repo "github.com/Jhon-2801/task-manager/core/repo"
	"github.com/Jhon-2801/task-manager/core/service"
	"github.com/Jhon-2801/task-manager/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectionBD()

	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := repo.NewRepo(db)
	userSrv := service.NewService(userRepo)
	userEnd := con.MakeEnponints(userSrv)

	router := gin.Default()

	router.POST("/login")
	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))

	router.Run("localhost:8080")
}
