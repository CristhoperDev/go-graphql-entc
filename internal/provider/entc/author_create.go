// Code generated by entc, DO NOT EDIT.

package entc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/author"
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/film"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// AuthorCreate is the builder for creating a Author entity.
type AuthorCreate struct {
	config
	mutation *AuthorMutation
	hooks    []Hook
}

// SetStatus sets the status field.
func (ac *AuthorCreate) SetStatus(i int8) *AuthorCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the status field if the given value is not nil.
func (ac *AuthorCreate) SetNillableStatus(i *int8) *AuthorCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// SetCreatedAt sets the created_at field.
func (ac *AuthorCreate) SetCreatedAt(t time.Time) *AuthorCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ac *AuthorCreate) SetNillableCreatedAt(t *time.Time) *AuthorCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the updated_at field.
func (ac *AuthorCreate) SetUpdatedAt(t time.Time) *AuthorCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (ac *AuthorCreate) SetNillableUpdatedAt(t *time.Time) *AuthorCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetFullName sets the fullName field.
func (ac *AuthorCreate) SetFullName(s string) *AuthorCreate {
	ac.mutation.SetFullName(s)
	return ac
}

// SetAge sets the age field.
func (ac *AuthorCreate) SetAge(i int8) *AuthorCreate {
	ac.mutation.SetAge(i)
	return ac
}

// SetID sets the id field.
func (ac *AuthorCreate) SetID(u uuid.UUID) *AuthorCreate {
	ac.mutation.SetID(u)
	return ac
}

// AddFilmIDs adds the films edge to Film by ids.
func (ac *AuthorCreate) AddFilmIDs(ids ...uuid.UUID) *AuthorCreate {
	ac.mutation.AddFilmIDs(ids...)
	return ac
}

// AddFilms adds the films edges to Film.
func (ac *AuthorCreate) AddFilms(f ...*Film) *AuthorCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ac.AddFilmIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (ac *AuthorCreate) Mutation() *AuthorMutation {
	return ac.mutation
}

// Save creates the Author in the database.
func (ac *AuthorCreate) Save(ctx context.Context) (*Author, error) {
	if _, ok := ac.mutation.Status(); !ok {
		v := author.DefaultStatus
		ac.mutation.SetStatus(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := author.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := author.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.FullName(); !ok {
		return nil, &ValidationError{Name: "fullName", err: errors.New("entc: missing required field \"fullName\"")}
	}
	if _, ok := ac.mutation.Age(); !ok {
		return nil, &ValidationError{Name: "age", err: errors.New("entc: missing required field \"age\"")}
	}
	var (
		err  error
		node *Author
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthorCreate) SaveX(ctx context.Context) *Author {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AuthorCreate) sqlSave(ctx context.Context) (*Author, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}

func (ac *AuthorCreate) createSpec() (*Author, *sqlgraph.CreateSpec) {
	var (
		a     = &Author{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: author.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: author.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		a.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldStatus,
		})
		a.Status = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: author.FieldCreatedAt,
		})
		a.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: author.FieldUpdatedAt,
		})
		a.UpdatedAt = value
	}
	if value, ok := ac.mutation.FullName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: author.FieldFullName,
		})
		a.FullName = value
	}
	if value, ok := ac.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldAge,
		})
		a.Age = value
	}
	if nodes := ac.mutation.FilmsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   author.FilmsTable,
			Columns: author.FilmsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: film.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}
