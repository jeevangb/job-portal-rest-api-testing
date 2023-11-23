package database

import (
	"jeevan/jobportal/config"

	"github.com/go-redis/redis/v8"
)

func ResdisConnection(cfg config.RedisConfig) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,     //address
		Password: cfg.RedisPassword, // No password by default
		DB:       cfg.RedisDb,       // Default DB
	})
	return redisClient
}
