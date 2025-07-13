package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisURL      string

	ServiceName string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println("error!!!", err)
	}
	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "postgres"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "akromjonotaboyev"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))
	cfg.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "food_api_gateway"))

	cfg.RedisURL = cast.ToString(getOrReturnDefault("REDIS_URL", "redis://default:ASw0AAIjcDFjNWExZjZiYzNiZmI0ZDYyYmQ1YjZkZmQxN2UwYzM3ZXAxMA@direct-muskox-11316.upstash.io:6379"))
	cfg.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "direct-muskox-11316.upstash.io"))
	cfg.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", "6379"))
	cfg.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "ASw0AAIjcDFjNWExZjZiYzNiZmI0ZDYyYmQ1YjZkZmQxN2UwYzM3ZXAxMA"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {

	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
