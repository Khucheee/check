package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Answer struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/health", checkHealth)
	router.Run("localhost:8080")
}

func checkHealth(c *gin.Context) {
	answer := Answer{"UP"}
	c.IndentedJSON(http.StatusOK, answer)
}
