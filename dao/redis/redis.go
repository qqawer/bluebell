package redis

import (
	"WebApp/global"
	"fmt"

	"github.com/go-redis/redis"
)

func Init() (err error) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     global.AppConfig.Redis.Host + ":" + fmt.Sprintf("%d", global.AppConfig.Redis.Port),
		DB:       global.AppConfig.Redis.DB,
		Password: global.AppConfig.Redis.Password,
		PoolSize: global.AppConfig.Redis.PoolSize,
	})
	_, err = RedisClient.Ping().Result()
	global.RedisDB = RedisClient
	return 

}
