package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// DetectLanguage detects the language from the request header
func DetectLanguage() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			language := "fa"
			val := c.Request().Header.Get("Accept-Language")
			if strings.Contains(val, "en") {
				language = "en"
			}
			c.Set("language", language)
			return next(c)
		}
	}
}
