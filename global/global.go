package global

import (
	"WebApp/settings"
	"github.com/go-redis/redis"
	
	"gorm.io/gorm"
)

var (
	Db      *gorm.DB
	RedisDB *redis.Client
	AppConfig *settings.Config
)