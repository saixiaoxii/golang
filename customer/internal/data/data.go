package data

import (
	"customer/internal/conf"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCustomerData)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	//初始化Rdb
	redisURL := fmt.Sprintf("redis://%s/1?dial_timeout=%d", c.Redis.Addr, 5)
	data := &Data{}
	options, _ := redis.ParseURL(redisURL)
	data.Rdb = redis.NewClient(options)

	cleanup := func() {
		_ = data.Rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil // 返回已经初始化的 data 实例
}
