package redis

import (
	"WebApp/global"
	"WebApp/settings"
	"fmt"

	"github.com/go-redis/redis"
)

func Init(cfg *settings.Config) (err error) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + fmt.Sprintf("%d", cfg.Redis.Port),
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
		PoolSize: cfg.Redis.PoolSize,
	})
	_, err = RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	global.RedisDB = RedisClient
	return nil

}
