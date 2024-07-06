package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luciorim/proxy-server/internal/service"
	"net/http"
)

type ProxyController struct {
	ProxyService *service.ProxyService
}

func NewProxyController(service *service.ProxyService) *ProxyController {
	return &ProxyController{
		ProxyService: service,
	}
}

func (p *ProxyController) InitRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/proxy", p.ProxyService.ProcessProxy)
	r.GET("/proxy/:id", p.ProxyService.GetById)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r

}
