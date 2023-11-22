package database

import "github.com/go-redis/redis/v8"

func ResdisConnection() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})
	return redisClient
}
