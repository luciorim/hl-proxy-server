package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luciorim/proxy-server/internal/cache"
	"github.com/luciorim/proxy-server/internal/dto"
	"io"
	"net/http"
)

type ProxyService struct {
	Cache *cache.Cache
}

func NewProxyService(proxyCache *cache.Cache) *ProxyService {
	return &ProxyService{
		Cache: proxyCache,
	}
}

var id int

func (p *ProxyService) ProcessProxy(ctx *gin.Context) {
	var proxyReq dto.ProxyRequest

	if err := ctx.ShouldBind(&proxyReq); err != nil {
		saveError(p.Cache, ctx, proxyReq, http.StatusBadRequest, err.Error())
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(*proxyReq.Method, *proxyReq.URL, nil)
	if err != nil {
		saveError(p.Cache, ctx, proxyReq, http.StatusBadRequest, err.Error())
		return
	}

	for k, v := range *proxyReq.Headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		saveError(p.Cache, ctx, proxyReq, http.StatusBadGateway, err.Error())

	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		saveError(p.Cache, ctx, proxyReq, http.StatusBadRequest, err.Error())
		return
	}

	proxyRes := dto.ProxyResponse{
		ID:      fmt.Sprintf("%d", generateId()),
		URL:     *proxyReq.URL,
		Status:  res.StatusCode,
		Headers: res.Header,
		Length:  len(resBody),
	}

	p.Cache.Set(&proxyReq, &proxyRes)

	ctx.JSON(http.StatusOK, proxyRes)

	defer res.Body.Close()

	return
}

func (p *ProxyService) GetById(ctx *gin.Context) {

	idFromPath := ctx.Param("id")

	req, res := p.Cache.Get(idFromPath)

	if res == nil || req == nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Request with ID: %d not found", id))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       id,
		"request":  req,
		"response": res,
	})

	return

}

func saveError(c *cache.Cache, ctx *gin.Context, proxyReq dto.ProxyRequest, statusCode int, errorMessage string) {
	errResponseMap := make(map[string][]string)
	errResponseMap["error"] = []string{errorMessage}

	proxyRes := dto.ProxyResponse{
		ID:      fmt.Sprintf("%d", generateId()),
		URL:     *proxyReq.URL,
		Status:  statusCode,
		Headers: errResponseMap,
		Length:  0,
	}

	c.Set(&proxyReq, &proxyRes)

	ctx.JSON(http.StatusOK, proxyRes)

}

func generateId() int {
	id++
	return id
}
