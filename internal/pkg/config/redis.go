package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitialiseRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Redis'in bağlı olup olmadığını kontrol etme
	pong, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis Connection Failed",
			err)
	}

	log.Println("Redis Successfully Connected.",
		"Ping", pong)

	RedisClient = conn

	return RedisClient
}
