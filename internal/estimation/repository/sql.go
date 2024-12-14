package repository

import (
	"context"
	"time"

	"github.com/Rustixir/go-challenge/ent"
	"github.com/Rustixir/go-challenge/ent/usersegment"
	"github.com/Rustixir/go-challenge/internal/estimation/entities"
)

type SQLRepository struct {
	db *ent.UserSegmentClient
}

func NewSQLRepository(db *ent.Client) *SQLRepository {
	return &SQLRepository{db: db.UserSegment}
}

func (s *SQLRepository) Create(ctx context.Context, entity entities.CreateRequest) error {
	_, err := s.db.Create().
		SetUserID(entity.UserID).
		SetSegment(entity.Segment).
		Save(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s *SQLRepository) Count(ctx context.Context, entity entities.CountRequest) (entities.CountResponse, error) {
	twoWeeksAgo := time.Now().AddDate(0, 0, -14)
	count, err := s.db.Query().Where(
		usersegment.SegmentEQ(entity.Segment),
		usersegment.CreatedAtGTE(twoWeeksAgo),
	).Count(ctx)
	if err != nil {
		return entities.CountResponse{}, err
	}
	return entities.CountResponse{
		Count: int64(count),
	}, nil
}

func (s *SQLRepository) Cleanup(ctx context.Context) error {
	twoWeeksAgo := time.Now().AddDate(0, 0, -14)
	_, err := s.db.Delete().
		Where(
			usersegment.CreatedAtLT(twoWeeksAgo),
		).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
