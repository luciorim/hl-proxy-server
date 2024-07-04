package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luciorim/proxy-server/internal/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.POST("/proxy", handler.ProxyRequest)
	r.GET("/proxy", handler.GetAllProxyRequests)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}

	log.Printf("Listening on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
