package dtos

import "github.com/Rustixir/go-challenge/pkg/errutil"

type CreateRequest struct {
	UserID  string
	Segment string
}

func (c CreateRequest) Validate() error {
	if c.UserID == "" {
		return errutil.ErrInvalidUser
	}
	if c.Segment == "" {
		return errutil.ErrInvalidSegment
	}
	return nil
}
