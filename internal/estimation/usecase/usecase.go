package usecase

import (
	"context"

	"github.com/Rustixir/go-challenge/internal/estimation/dtos"
)

type Usecase interface {
	CreateUserSegment(ctx context.Context, model dtos.CreateRequest) error
	EstimateSegment(ctx context.Context, model dtos.CountRequest) (dtos.CountResponse, error)
}
