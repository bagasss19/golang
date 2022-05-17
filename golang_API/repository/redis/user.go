package redisrepository

import (
	"context"
	"fmt"
	"golang_api/model"

	"github.com/go-redis/redis/v8"
)

type UserRedisRepository interface {
	GetAllUserRedis(ctx context.Context) (data string, err error)
	SetAllUserRedis(ctx context.Context) (data []*model.User, err error)
	DelAllUserRedis(ctx context.Context) error
}

func NewUserRedisRepository(redisConn *redis.Client) UserRedisRepository {
	return &UserRedisRepositoryDB{
		repoPrefix: "userRedisRepo",
		redisConn:  redisConn,
	}
}

func (u *UserRedisRepositoryDB) SetAllUserRedis(ctx context.Context) (data []*model.User, err error) {
	prefix := u.repoPrefix + ".SetAllUserRedis"
	query := u.redisConn.WithContext(ctx)

	if err := query.Set(ctx, "Users", data, 0).Err(); err != nil {
		return nil, fmt.Errorf("[%s] %+v", prefix, err)
	}
	return data, nil
}

func (u *UserRedisRepositoryDB) GetAllUserRedis(ctx context.Context) (data string, err error) {
	prefix := u.repoPrefix + ".GetAllUserRedis"
	query := u.redisConn.WithContext(ctx)

	result, err := query.Get(ctx, "Users").Result()
	if err != nil {
		return "", fmt.Errorf("[%s] %+v", prefix, err)
	}

	return result, nil
}

func (u *UserRedisRepositoryDB) DelAllUserRedis(ctx context.Context) error {
	prefix := u.repoPrefix + ".DelAllUserRedis"
	query := u.redisConn.WithContext(ctx)

	if err := query.Del(ctx, "Users").Err(); err != nil {
		return fmt.Errorf("[%s] %+v", prefix, err)
	}
	return nil
}
