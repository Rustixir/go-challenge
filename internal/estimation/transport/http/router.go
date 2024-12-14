package http

import "github.com/labstack/echo/v4"

func (h *Handlers) Register(domain *echo.Group) {
	domain.POST("/segment", h.CreateUserSegment)
	domain.GET("/segment/:name/count", h.GetEstimateSegment)
}
