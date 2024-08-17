package connector

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

func CreateRedisClient() (*redis.Client, error) {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		DB:   db,
	})

	_, err = client.Ping().Result()

	if err != nil {
		return nil, err
	}

	return client, nil
}
