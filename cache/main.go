package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"singo/config"
	"singo/util"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// Redis 在中间件中初始化redis链接
func Redis(ctx context.Context) {
	cfg := config.GetConfig()

	client := redis.NewClient(&redis.Options{
		Addr:       cfg.GetString("redis.ip"),
		Password:   cfg.GetString("redis.pwd"),
		DB:         cfg.GetInt("redis.db"),
		MaxRetries: 1,
	})

	_, err := client.Ping(ctx).Result()

	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}
