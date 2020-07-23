// Code generated by entc, DO NOT EDIT.

package entc

import (
	"context"
	"fmt"

	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/film"
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FilmDelete is the builder for deleting a Film entity.
type FilmDelete struct {
	config
	hooks      []Hook
	mutation   *FilmMutation
	predicates []predicate.Film
}

// Where adds a new predicate to the delete builder.
func (fd *FilmDelete) Where(ps ...predicate.Film) *FilmDelete {
	fd.predicates = append(fd.predicates, ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FilmDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FilmMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FilmDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FilmDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: film.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: film.FieldID,
			},
		},
	}
	if ps := fd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FilmDeleteOne is the builder for deleting a single Film entity.
type FilmDeleteOne struct {
	fd *FilmDelete
}

// Exec executes the deletion query.
func (fdo *FilmDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{film.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FilmDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}