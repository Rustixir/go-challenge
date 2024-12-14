// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Rustixir/go-challenge/ent/schema"
	"github.com/Rustixir/go-challenge/ent/usersegment"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersegmentFields := schema.UserSegment{}.Fields()
	_ = usersegmentFields
	// usersegmentDescCreatedAt is the schema descriptor for created_at field.
	usersegmentDescCreatedAt := usersegmentFields[2].Descriptor()
	// usersegment.DefaultCreatedAt holds the default value on creation for the created_at field.
	usersegment.DefaultCreatedAt = usersegmentDescCreatedAt.Default.(func() time.Time)
}
