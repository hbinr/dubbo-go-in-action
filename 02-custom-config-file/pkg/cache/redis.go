package cache

import (
	"fmt"

	"study.dubbogo/02-custom-config-file/pkg/conf"

	"github.com/apache/dubbo-go/common/logger"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init init redis client
func InitRedis(cfg *conf.DataConfig) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.RedisConfig.Host, cfg.RedisConfig.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default db
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
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
