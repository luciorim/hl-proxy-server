package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/luciorim/proxy-server/internal/handler"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	r := gin.Default()
	r.POST("/proxy", handler.ProxyRequest)
	r.GET("/proxy", handler.GetAllProxyRequests)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	log.Printf("Server started on port %s\n", port)

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server startup failed: %v\n", err)
	}
}
