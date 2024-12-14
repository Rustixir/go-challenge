package dtos

import "github.com/Rustixir/go-challenge/pkg/errutil"

type CountRequest struct {
	Segment string
}

func (c CountRequest) Validate() error {
	if c.Segment == "" {
		return errutil.ErrInvalidSegment
	}
	return nil
}

type CountResponse struct {
	Count int64
}
