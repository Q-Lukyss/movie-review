package main

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Let's have a Gin and Tonic!"})
    })

    r.Run(":8080") // Écoute sur le port 8080
}
