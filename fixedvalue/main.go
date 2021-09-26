package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Handle("GET", "fixedvalue", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"value": "actualvalue"})
	})

	engine.Run()
}
