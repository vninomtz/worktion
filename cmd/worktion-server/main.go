package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vninomtz/worktion/pkg/db"
	"github.com/vninomtz/worktion/pkg/handler"
	"github.com/vninomtz/worktion/pkg/repository"
	"github.com/vninomtz/worktion/pkg/service"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal("Unable to initialize db", err)
	}

	userRepo := repository.NewSqliteExerciseRepo(db)

	userService := service.NewExerciseService(userRepo)

	router := gin.Default()

	handler.NewHandler(router, userService)

	router.Run()
}

func hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"response": "hello"})
}
