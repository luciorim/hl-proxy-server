package cache

import (
	"github.com/luciorim/proxy-server/internal/dto"
	"github.com/patrickmn/go-cache"
)

type Cache struct {
	Cache *cache.Cache
}

func NewCache() *Cache {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)

	return &Cache{
		Cache: c,
	}
}

func (c *Cache) Set(req *dto.ProxyRequest, res *dto.ProxyResponse) {

	id := res.ID

	c.Cache.Set(id+"_req", req, cache.NoExpiration)
	c.Cache.Set(id+"_res", res, cache.NoExpiration)

}

func (c *Cache) Get(id string) (interface{}, interface{}) {

	reqId := id + "_req"
	resId := id + "_res"

	reqData, ok1 := c.Cache.Get(reqId)
	resData, ok2 := c.Cache.Get(resId)

	if !ok1 || !ok2 {
		return nil, nil
	}

	return reqData.(*dto.ProxyRequest), resData.(*dto.ProxyResponse)
}
