package repository

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/Rustixir/go-challenge/internal/estimation/entities"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestCleanup(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	err := client.FlushDB(context.Background()).Err()
	assert.NoError(t, err, "failed to flush Redis before test")

	repo := NewRedisRepository(client)

	ctx := context.Background()
	now := time.Now()
	segments := []string{"sports", "technology", "health"}
	expiredCount, validCount := 0, 0

	for _, segment := range segments {
		for i := 0; i < 10; i++ {
			userID := fmt.Sprintf("user_%d", rand.Intn(1000))
			daysAgo := rand.Intn(30)
			userTime := now.Add(time.Duration(-daysAgo*24) * time.Hour)
			err := repo.Create(ctx, entities.CreateRequest{
				UserID:  userID,
				Segment: segment,
			})
			assert.NoError(t, err, "failed to create user")
			if daysAgo > 14 {
				expiredCount++
			} else {
				validCount++
			}
			client.ZAdd(ctx, fmt.Sprintf("segment:%s", segment), redis.Z{
				Score:  float64(userTime.Unix()),
				Member: userID,
			})
		}
	}
	err = repo.Cleanup(ctx)
	assert.NoError(t, err, "cleanup failed")
	for _, segment := range segments {
		response, err := repo.Count(ctx, entities.CountRequest{Segment: segment})
		assert.NoError(t, err, "failed to count users in segment")
		assert.LessOrEqual(t, response.Count, int64(validCount),
			fmt.Sprintf("segment %s has more users than expected after cleanup", segment))
	}
	t.Logf("Expired users removed: %d", expiredCount)
	t.Logf("Valid users remaining: %d", validCount)
}
