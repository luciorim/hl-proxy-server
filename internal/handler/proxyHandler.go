package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luciorim/proxy-server/internal/dto"
	"io"
	"net/http"
	"sync"
)

var (
	db sync.Map
	id int
)

func ProxyRequest(ctx *gin.Context) {

	var proxyReq dto.ProxyRequest

	if err := ctx.ShouldBind(&proxyReq); err != nil {
		saveError(ctx, proxyReq, http.StatusBadRequest, err.Error())
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(*proxyReq.Method, *proxyReq.URL, nil)
	if err != nil {
		saveError(ctx, proxyReq, http.StatusBadRequest, err.Error())
		return
	}

	for k, v := range *proxyReq.Headers {
		req.Header.Add(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		saveError(ctx, proxyReq, http.StatusBadGateway, err.Error())

	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		saveError(ctx, proxyReq, http.StatusBadRequest, err.Error())
		return
	}

	proxyRes := dto.ProxyResponse{
		ID:      fmt.Sprintf("%d", generateId()),
		URL:     *proxyReq.URL,
		Status:  res.StatusCode,
		Headers: res.Header,
		Length:  len(resBody),
	}

	db.Store(proxyRes.ID+"_req", proxyReq)
	db.Store(proxyRes.ID+"_res", proxyRes)

	ctx.JSON(http.StatusOK, proxyRes)

	defer res.Body.Close()

}

func GetAllProxyRequests(ctx *gin.Context) {
	requests := make(map[string]interface{})

	db.Range(func(k, v interface{}) bool {
		requests[k.(string)] = v
		return true
	})

	ctx.JSON(http.StatusOK, requests)
}

func generateId() int {
	id++
	return id
}

func saveError(ctx *gin.Context, proxyReq dto.ProxyRequest, statusCode int, errorMessage string) {
	errResponseMap := make(map[string][]string)
	errResponseMap["error"] = []string{errorMessage}

	proxyRes := dto.ProxyResponse{
		ID:      fmt.Sprintf("%d", generateId()),
		URL:     *proxyReq.URL,
		Status:  statusCode,
		Headers: errResponseMap,
		Length:  0,
	}

	db.Store(proxyRes.ID+"_req", proxyReq)
	db.Store(proxyRes.ID+"_res", proxyRes)

	ctx.JSON(http.StatusOK, proxyRes)

}
