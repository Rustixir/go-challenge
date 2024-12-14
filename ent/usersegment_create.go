// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Rustixir/go-challenge/ent/usersegment"
)

// UserSegmentCreate is the builder for creating a UserSegment entity.
type UserSegmentCreate struct {
	config
	mutation *UserSegmentMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (usc *UserSegmentCreate) SetUserID(s string) *UserSegmentCreate {
	usc.mutation.SetUserID(s)
	return usc
}

// SetSegment sets the "segment" field.
func (usc *UserSegmentCreate) SetSegment(s string) *UserSegmentCreate {
	usc.mutation.SetSegment(s)
	return usc
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserSegmentCreate) SetCreatedAt(t time.Time) *UserSegmentCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserSegmentCreate) SetNillableCreatedAt(t *time.Time) *UserSegmentCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// Mutation returns the UserSegmentMutation object of the builder.
func (usc *UserSegmentCreate) Mutation() *UserSegmentMutation {
	return usc.mutation
}

// Save creates the UserSegment in the database.
func (usc *UserSegmentCreate) Save(ctx context.Context) (*UserSegment, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSegmentCreate) SaveX(ctx context.Context) *UserSegment {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSegmentCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSegmentCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSegmentCreate) defaults() {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		v := usersegment.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSegmentCreate) check() error {
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserSegment.user_id"`)}
	}
	if _, ok := usc.mutation.Segment(); !ok {
		return &ValidationError{Name: "segment", err: errors.New(`ent: missing required field "UserSegment.segment"`)}
	}
	if _, ok := usc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserSegment.created_at"`)}
	}
	return nil
}

func (usc *UserSegmentCreate) sqlSave(ctx context.Context) (*UserSegment, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSegmentCreate) createSpec() (*UserSegment, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSegment{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersegment.Table, sqlgraph.NewFieldSpec(usersegment.FieldID, field.TypeInt))
	)
	if value, ok := usc.mutation.UserID(); ok {
		_spec.SetField(usersegment.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := usc.mutation.Segment(); ok {
		_spec.SetField(usersegment.FieldSegment, field.TypeString, value)
		_node.Segment = value
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.SetField(usersegment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// UserSegmentCreateBulk is the builder for creating many UserSegment entities in bulk.
type UserSegmentCreateBulk struct {
	config
	err      error
	builders []*UserSegmentCreate
}

// Save creates the UserSegment entities in the database.
func (uscb *UserSegmentCreateBulk) Save(ctx context.Context) ([]*UserSegment, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSegment, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSegmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSegmentCreateBulk) SaveX(ctx context.Context) []*UserSegment {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSegmentCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSegmentCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}