package redis

import (
	"backend_course/rent_car/config"
	"backend_course/rent_car/storage"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	db *redis.Client
}

func New(cfg config.Config) storage.IRedisStorage {
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
	})

	return Store{
		db: redisClient,
	}
}

func (s Store) SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	statusCmd := s.db.SetEx(ctx, key, value, duration)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}

	return nil
}

func (s Store) Get(ctx context.Context, key string) (interface{}, error) {
	resp := s.db.Get(ctx, key)

	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return resp.Val(), nil
}

func (s Store) Del(ctx context.Context, key string) error {
	statusCmd := s.db.Del(ctx, key)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}
