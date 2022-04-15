package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client = nil

func GetRedisClient() (*redis.Client, error) {
	if rdb != nil {
		return rdb, nil
	}
	opts := &redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	}
	rdb = redis.NewClient(opts)
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return rdb, err
}
