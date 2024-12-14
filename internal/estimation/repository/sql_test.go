package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	"github.com/Rustixir/go-challenge/ent/usersegment"
	"github.com/Rustixir/go-challenge/internal/estimation/entities"
	_ "github.com/mattn/go-sqlite3"

	"github.com/Rustixir/go-challenge/ent"
	"github.com/stretchr/testify/assert"
)

// WARNING: should export CGO_ENABLED=1
func setupTestDB(t *testing.T) *ent.Client {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return client
}

func TestSQLRepository(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()
	repo := NewSQLRepository(client)
	ctx := context.Background()
	segments := []string{"sports", "technology", "health"}
	now := time.Now()
	expiredUsers, validUsers := 0, 0
	for _, segment := range segments {
		for i := 0; i < 10; i++ {
			userID := fmt.Sprintf("user_%d_%s", i, segment)
			createdAt := now.Add(time.Duration(-i*3*24) * time.Hour)
			_, err := client.UserSegment.Create().
				SetUserID(userID).
				SetSegment(segment).
				SetCreatedAt(createdAt).
				Save(ctx)
			assert.NoError(t, err, "failed to create user in segment")

			if i >= 5 {
				expiredUsers++
			} else {
				validUsers++
			}
		}
	}
	for _, segment := range segments {
		response, err := repo.Count(ctx, entities.CountRequest{Segment: segment})
		assert.NoError(t, err, "failed to count users for segment")
		assert.Equal(t, int64(validUsers/3), response.Count, "user count mismatch")
	}

	err := repo.Cleanup(ctx)
	assert.NoError(t, err, "cleanup failed")
	expiredCount := 0
	for _, segment := range []string{"sports", "technology", "health"} {
		count, err := client.UserSegment.Query().
			Where(
				usersegment.SegmentEQ(segment),
				usersegment.CreatedAtLT(time.Now().Add(-14*24*time.Hour)),
			).Count(ctx)
		assert.NoError(t, err, "failed to query after cleanup")
		expiredCount += count
	}
	assert.Equal(t, 0, expiredCount, "expired users not cleaned up")
}
