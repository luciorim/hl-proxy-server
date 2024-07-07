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

// ProcessProxy godoc
//
//	@Summary 	process requested url
//	@Tags 		proxy-controller
//	@Description allows make http request to another website
//	@Accept 		json
//	@Produce 	json
//	@Param 		input  body     dto.ProxyRequest  true "request data"
//	@Success 	200    {object} dto.ProxyResponse
//	@Failure 	400    {object} string
//	@Failure 	502    {object} string
//	@Router 		/proxy [post]
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

// GetById godoc
//
//	@Summary get proxy history by id
//	@Tags proxy-controller
//	@Description return request and response by id
//	@Produce json
//	@Param id path int true "request id"
//	@Failure 400 {object} dto.ProxyResponse
//	@Router /proxy/{id} [get]
func (p *ProxyService) GetById(ctx *gin.Context) {

	idFromPath := ctx.Param("id")

	req, res := p.Cache.Get(idFromPath)

	if res == nil || req == nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Request with ID: %s not found", idFromPath))
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

	ctx.JSON(http.StatusBadRequest, proxyRes)

}

func generateId() int {
	id++
	return id
}
