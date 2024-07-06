package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/luciorim/proxy-server/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/luciorim/proxy-server/docs"
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}
