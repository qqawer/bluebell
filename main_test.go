package main

import (
	"WebApp/settings"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
)

// 测试 Redis 连接
func TestRedisConnection(t *testing.T) {
	// Initialize configuration
	err := settings.Init()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}

	// Create Redis client
	cfg := settings.AppConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
		PoolSize: cfg.Redis.PoolSize,
	})

	// Test the Redis connection
	_, err = rdb.Ping().Result()
	if err != nil {
		t.Fatalf("Redis connection failed: %v", err)
	}

	t.Log("Redis connection successful!")
}
