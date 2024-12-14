package config

import (
	"context"
	"log/slog"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var Config Schema

var once sync.Once

type Schema struct {
	Env  string `env:"APP_ENV"`
	Port int    `env:"APP_PORT,default=8086"`
	DB   struct {
		ActiveDatabase string `env:"APP_DB_ACTIVE_DATABASE" envDefault:"sqlite"`
		Redis          struct {
			Addr string `env:"APP_REDIS_ADDR,default=localhost:6379"`
		}
		SQL struct {
			DSN string `env:"APP_SQL_DSN,default=file:ent?mode=memory&cache=shared&_fk=1"`
		}
	}
}

func (s *Schema) ActiveDB() string {
	switch s.DB.ActiveDatabase {
	case "sqlite", "redis":
		return s.DB.ActiveDatabase
	default:
		return "sqlite"
	}
}

func init() {
	load()
}

func load() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			slog.Error("failed to load .env file", "error", err)
		}
		if err := envconfig.Process(context.Background(), &Config); err != nil {
			slog.Error("failed to load config", "error", err)
		}
	})
}
