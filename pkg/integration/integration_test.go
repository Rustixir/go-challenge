package integration

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Rustixir/go-challenge/internal"
	"github.com/Rustixir/go-challenge/internal/estimation/dtos"
)

func TestErrorHandler(t *testing.T) {
	go func() {
		internal.StartApp()
	}()
	time.Sleep(1 * time.Second)
	err := SendRequest("POST", "/segment", dtos.CreateRequest{}, nil)
	if err != nil {
		if strings.Contains(err.Error(), "کاربر نامعتبر است") {
			return
		}
	}
	t.Error("unexpected response")
}

func TestEstimationSegment(t *testing.T) {
	go func() {
		internal.StartApp()
	}()
	time.Sleep(1 * time.Second)
	segments := []string{"sports", "technology", "health"}
	for _, segment := range segments {
		for i := 0; i < 10; i++ {
			err := SendRequest("POST", "/segment", dtos.CreateRequest{
				UserID:  fmt.Sprintf("user-%v", time.Now().Unix()%1000),
				Segment: segment,
			}, nil)
			if err != nil {
				t.Errorf("failed to send request: %s", err)
			}
		}
	}
	for _, segment := range segments {
		var response dtos.CountResponse
		err := SendRequest("GET", fmt.Sprintf("/segment/%s/count", segment), nil, &response)
		if err != nil {
			t.Errorf("failed to send request: %s", err)
		}
		if response.Count != 10 {
			t.Errorf("count should be greater than 0")
		}
	}
	return
}
