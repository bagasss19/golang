package config

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var (
	_   = godotenv.Load(".secret.env")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
)
