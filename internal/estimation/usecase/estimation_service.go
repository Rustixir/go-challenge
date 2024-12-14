package usecase

import (
	"context"
	"log/slog"
	"time"

	"github.com/Rustixir/go-challenge/internal/estimation/dtos"
	"github.com/Rustixir/go-challenge/internal/estimation/entities"
	"github.com/Rustixir/go-challenge/internal/estimation/repository"
	"github.com/Rustixir/go-challenge/pkg/errutil"
)

type estimationService struct {
	repo repository.Repository
}

func NewEstimationService(repo repository.Repository) Usecase {
	svc := &estimationService{
		repo: repo,
	}
	go svc.cleanup()
	return svc
}

func (s *estimationService) CreateUserSegment(ctx context.Context, dto dtos.CreateRequest) error {
	err := s.repo.Create(ctx, entities.CreateRequest{
		UserID:  dto.UserID,
		Segment: dto.Segment,
	})
	if err != nil {
		slog.Error("failed to create user segment", "error", err)
		return errutil.ErrInternalServerError
	}
	return nil
}

func (s *estimationService) EstimateSegment(ctx context.Context, dto dtos.CountRequest) (dtos.CountResponse, error) {
	entity, err := s.repo.Count(ctx, entities.CountRequest{
		Segment: dto.Segment,
	})
	if err != nil {
		slog.Error("failed to count user segment", "error", err)
		return dtos.CountResponse{}, errutil.ErrInternalServerError
	}
	return dtos.CountResponse{
		Count: entity.Count,
	}, nil
}

func (s *estimationService) cleanup() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx := context.Background()
			err := s.repo.Cleanup(ctx)
			if err != nil {
				slog.Error("failed to cleanup expired segments", "error", err)
			} else {
				slog.Info("Expired segments cleaned up")
			}
		}
	}
}
