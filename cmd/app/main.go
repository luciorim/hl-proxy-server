package main

import (
	c "github.com/luciorim/proxy-server/internal/cache"
	"github.com/luciorim/proxy-server/internal/config"
	"github.com/luciorim/proxy-server/internal/controller"
	"github.com/luciorim/proxy-server/internal/service"
	"github.com/luciorim/proxy-server/internal/utils/logger"
)

// @title Proxy Server
// @version 1.0
// @description Server allows do http requests to another websites

// @host localhost:8181
// @BasePath /
func main() {

	cfg := config.MustInit()

	logger := logger.InitLogger(cfg.Env)

	cache := c.NewCache()

	proxyService := service.NewProxyService(cache)

	proxyController := controller.NewProxyController(proxyService)

	r := proxyController.InitRoutes()

	logger.Info("Server started on port: " + cfg.HTTPServer.Port)

	if err := r.Run(":" + cfg.HTTPServer.Port); err != nil {
		logger.Error(err.Error())
	}

}
