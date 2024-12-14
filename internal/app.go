package internal

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/Rustixir/go-challenge/internal/estimation"
	"github.com/Rustixir/go-challenge/internal/estimation/transport/http/middleware"
	"github.com/Rustixir/go-challenge/pkg/config"
	"github.com/Rustixir/go-challenge/pkg/errutil"
	"github.com/Rustixir/go-challenge/pkg/localization"
	"github.com/labstack/echo/v4"
)

func StartApp() {
	srv := echo.New()
	RegisterErrorHandler(srv)
	RegisterMiddlewares(srv)

	router := srv.Group("/api")
	estimation.StartService(router)

	err := srv.Start(fmt.Sprintf(":%d", config.Config.Port))
	slog.Error("failed to start server", "error", err)
}

func RegisterErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = ErrorHandler
}

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(middleware.DetectLanguage())

}

func ErrorHandler(err error, c echo.Context) {
	var keyErr *errutil.KeyError
	if ok := errors.As(err, &keyErr); !ok {
		keyErr = errutil.ErrInternalServerError
	}
	language := c.Get("language").(string)
	value := localization.Get(language, keyErr.Key)
	_ = c.JSON(keyErr.StatusCode, map[string]string{
		"error": value,
	})
}
