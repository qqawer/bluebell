package redis

import (
	"WebApp/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func Init(cfg *settings.Config) (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
		PoolSize: cfg.Redis.PoolSize,
	})
	_, err = RedisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil

}
func Close() {
	_ = RedisClient.Close()
}
