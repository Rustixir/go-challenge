package repository

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/Rustixir/go-challenge/internal/estimation/entities"
	"github.com/redis/go-redis/v9"
)

var (
	segmentKeyPrefix = "segment:%s"
)

type RedisRepository struct {
	rdb *redis.Client
}

func NewRedisRepository(rdb *redis.Client) Repository {
	return &RedisRepository{rdb}
}

func (r *RedisRepository) Create(ctx context.Context, entity entities.CreateRequest) error {
	key := fmt.Sprintf(segmentKeyPrefix, entity.Segment)
	err := r.rdb.ZAdd(ctx, key, redis.Z{
		Score:  float64(time.Now().Add(14 * 24 * time.Hour).Unix()),
		Member: entity.UserID,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) Count(ctx context.Context, entity entities.CountRequest) (entities.CountResponse, error) {
	key := fmt.Sprintf(segmentKeyPrefix, entity.Segment)
	twoWeeksAgo := strconv.FormatInt(time.Now().AddDate(0, 0, -14).Unix(), 10)
	count, err := r.rdb.ZCount(ctx, key, twoWeeksAgo, "+inf").Result()
	if err != nil {
		return entities.CountResponse{}, err
	}
	return entities.CountResponse{
		Count: count,
	}, nil
}

func (r *RedisRepository) Cleanup(ctx context.Context) error {
	iter := r.rdb.Scan(ctx, 0, "segment:*", 10000).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		twoWeeksAgo := strconv.FormatInt(time.Now().AddDate(0, 0, -14).Unix(), 10)
		// Remove all users with expiration timestamps less than the twoWeeksAgo.
		err := r.rdb.ZRemRangeByScore(ctx, key, "-inf", twoWeeksAgo).Err()
		if err != nil {
			slog.Error("failed to cleanup expired segment members", "error", err, "segment", key)
		}
	}
	return nil
}
