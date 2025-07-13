package postgres

import (
	"backend_course/rent_car/config"
	"backend_course/rent_car/pkg/logger"
	"backend_course/rent_car/storage"
	"backend_course/rent_car/storage/redis"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Store struct {
	Pool   *pgxpool.Pool
	logger logger.ILogger
	cfg    config.Config
	redis  storage.IRedisStorage
}

func New(ctx context.Context, cfg config.Config, logger logger.ILogger, redis storage.IRedisStorage) (storage.IStorage, error) {
	url := fmt.Sprintf(`host=%s port=%v user=%s password=%s database=%s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)

	pgPoolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	pgPoolConfig.MaxConns = 100
	pgPoolConfig.MaxConnLifetime = time.Hour

	newPool, err := pgxpool.NewWithConfig(context.Background(), pgPoolConfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	return Store{
		Pool:   newPool,
		logger: logger,
		cfg:    cfg,
		redis:  redis,
	}, nil
}

func (s Store) CloseDB() {
	s.Pool.Close()
}

func (s Store) Car() storage.ICarStorage {
	newCar := NewCarRepo(s.Pool, s.logger)

	return &newCar
}

func (s Store) Customer() storage.ICustomerStorage {
	newCustomer := NewCustomerRepo(s.Pool, s.logger)

	return &newCustomer
}

func (s Store) Order() storage.IOrderStorage {
	newOrder := NewOrderRepo(s.Pool, s.logger)

	return &newOrder
}

func (s Store) Redis() storage.IRedisStorage {
	return redis.New(s.cfg)
}
