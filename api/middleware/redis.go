package middleware

import (
	"github.com/go-redis/redis"
)

func Cache() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
    })
	return client
}