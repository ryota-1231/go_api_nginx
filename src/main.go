package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/api/greeting", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	engine.Run(":8000")
}
