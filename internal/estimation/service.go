package estimation

import (
	"github.com/Rustixir/go-challenge/internal/estimation/repository"
	"github.com/Rustixir/go-challenge/internal/estimation/transport/http"
	"github.com/Rustixir/go-challenge/internal/estimation/usecase"
	"github.com/Rustixir/go-challenge/internal/provider"
	"github.com/Rustixir/go-challenge/pkg/config"
	"github.com/labstack/echo/v4"
)

func StartService(router *echo.Group) {
	http.NewHandlers(
		usecase.NewEstimationService(
			repoFactory(),
		),
	).Register(router)
}

func repoFactory() repository.Repository {
	switch config.Config.ActiveDB() {
	case "redis":
		return repository.NewRedisRepository(
			provider.GetRedis(),
		)
	default:
		return repository.NewSQLRepository(
			provider.GetSqliteEnt(),
		)
	}
}
