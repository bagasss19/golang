package redisrepository

import "github.com/go-redis/redis/v8"

type UserRedisRepositoryDB struct {
	repoPrefix string
	redisConn  *redis.Client
}
