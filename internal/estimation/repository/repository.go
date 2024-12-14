package repository

import (
	"context"

	"github.com/Rustixir/go-challenge/internal/estimation/entities"
)

type Repository interface {
	Create(ctx context.Context, entity entities.CreateRequest) error
	Count(ctx context.Context, entity entities.CountRequest) (entities.CountResponse, error)
	Cleanup(ctx context.Context) error
}
