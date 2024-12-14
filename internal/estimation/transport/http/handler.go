package http

import (
	"net/http"

	"github.com/Rustixir/go-challenge/internal/estimation/dtos"
	"github.com/Rustixir/go-challenge/internal/estimation/usecase"
	"github.com/Rustixir/go-challenge/pkg/errutil"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	usecase usecase.Usecase
}

func NewHandlers(usecase usecase.Usecase) *Handlers {
	return &Handlers{
		usecase: usecase,
	}
}

func (h *Handlers) CreateUserSegment(ctx echo.Context) error {
	var request dtos.CreateRequest
	if err := ctx.Bind(&request); err != nil {
		return errutil.ErrInvalidRequest
	}
	if err := request.Validate(); err != nil {
		return err
	}
	if err := h.usecase.CreateUserSegment(ctx.Request().Context(), request); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (h *Handlers) GetEstimateSegment(ctx echo.Context) error {
	var request dtos.CountRequest
	request.Segment = ctx.Param("name")
	if err := request.Validate(); err != nil {
		return err
	}
	response, err := h.usecase.EstimateSegment(ctx.Request().Context(), request)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, response)
}
