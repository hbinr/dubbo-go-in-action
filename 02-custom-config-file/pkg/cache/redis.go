package cache

import (
	"fmt"

	"study.dubbogo/02-custom-config-file/pkg/conf"

	"github.com/apache/dubbo-go/common/logger"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init init redis client
func InitRedis(cfg *conf.AppConfig) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		Password:     cfg.Password,     // no password set
		DB:           cfg.DB,           // use default db
		PoolSize:     cfg.PoolSize,     // redis connection pool size
		MinIdleConns: cfg.MinIdleConns, // Set the minimum number of connections in the idle connection pool
	})
	if _, err := rdb.Ping().Result(); err != nil {
		logger.Error("redis ping failed", err)
		return nil, err
	}
	return rdb, nil
}

// Close close redis client
func Close() {
	_ = rdb.Close()
}
