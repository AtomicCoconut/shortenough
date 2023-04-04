package database

import (
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func Connect() {
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	client = rc
}
