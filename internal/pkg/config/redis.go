package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var RedisClient *redis.Client

func InitialiseRedis() *redis.Client {
	godotenv.Load()
	dsn := os.Getenv("REDIS_URL")

	conn := redis.NewClient(&redis.Options{
		Addr: dsn,
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
