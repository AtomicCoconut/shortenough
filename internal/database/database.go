package database

import (
	"context"
	"time"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()
var Nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func Connect() {
	rc := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	client = rc
}

func CreateShort(url string) (string, error) {
	short, err := uuid.NewV5(uuid.NamespaceURL, []byte(time.Now().String()))
	if err != nil {
		return "", err
	}
	shortString := short.String()

	err = client.Set(ctx, shortString, url, 0).Err()
	if err != nil {
		return "", err
	}
	return shortString, nil
}

func FindShort(short string) (string, error) {
	return client.Get(ctx, short).Result()
}
