package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("exercise", hello)
	}

	router.Run()
}

func hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"response": "hello"})
}
