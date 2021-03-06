// Code generated by entc, DO NOT EDIT.

package entc

import (
	"context"
	"fmt"
	"time"

	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/author"
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/film"
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// AuthorUpdate is the builder for updating Author entities.
type AuthorUpdate struct {
	config
	hooks      []Hook
	mutation   *AuthorMutation
	predicates []predicate.Author
}

// Where adds a new predicate for the builder.
func (au *AuthorUpdate) Where(ps ...predicate.Author) *AuthorUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetStatus sets the status field.
func (au *AuthorUpdate) SetStatus(i int8) *AuthorUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(i)
	return au
}

// SetNillableStatus sets the status field if the given value is not nil.
func (au *AuthorUpdate) SetNillableStatus(i *int8) *AuthorUpdate {
	if i != nil {
		au.SetStatus(*i)
	}
	return au
}

// AddStatus adds i to status.
func (au *AuthorUpdate) AddStatus(i int8) *AuthorUpdate {
	au.mutation.AddStatus(i)
	return au
}

// SetUpdatedAt sets the updated_at field.
func (au *AuthorUpdate) SetUpdatedAt(t time.Time) *AuthorUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetFullName sets the fullName field.
func (au *AuthorUpdate) SetFullName(s string) *AuthorUpdate {
	au.mutation.SetFullName(s)
	return au
}

// SetAge sets the age field.
func (au *AuthorUpdate) SetAge(i int8) *AuthorUpdate {
	au.mutation.ResetAge()
	au.mutation.SetAge(i)
	return au
}

// AddAge adds i to age.
func (au *AuthorUpdate) AddAge(i int8) *AuthorUpdate {
	au.mutation.AddAge(i)
	return au
}

// AddFilmIDs adds the films edge to Film by ids.
func (au *AuthorUpdate) AddFilmIDs(ids ...uuid.UUID) *AuthorUpdate {
	au.mutation.AddFilmIDs(ids...)
	return au
}

// AddFilms adds the films edges to Film.
func (au *AuthorUpdate) AddFilms(f ...*Film) *AuthorUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.AddFilmIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (au *AuthorUpdate) Mutation() *AuthorMutation {
	return au.mutation
}

// RemoveFilmIDs removes the films edge to Film by ids.
func (au *AuthorUpdate) RemoveFilmIDs(ids ...uuid.UUID) *AuthorUpdate {
	au.mutation.RemoveFilmIDs(ids...)
	return au
}

// RemoveFilms removes films edges to Film.
func (au *AuthorUpdate) RemoveFilms(f ...*Film) *AuthorUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.RemoveFilmIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AuthorUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := author.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthorUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthorUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthorUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AuthorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   author.Table,
			Columns: author.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: author.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldStatus,
		})
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldStatus,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: author.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.FullName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: author.FieldFullName,
		})
	}
	if value, ok := au.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldAge,
		})
	}
	if value, ok := au.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldAge,
		})
	}
	if nodes := au.mutation.RemovedFilmsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.FilmsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{author.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AuthorUpdateOne is the builder for updating a single Author entity.
type AuthorUpdateOne struct {
	config
	hooks    []Hook
	mutation *AuthorMutation
}

// SetStatus sets the status field.
func (auo *AuthorUpdateOne) SetStatus(i int8) *AuthorUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(i)
	return auo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (auo *AuthorUpdateOne) SetNillableStatus(i *int8) *AuthorUpdateOne {
	if i != nil {
		auo.SetStatus(*i)
	}
	return auo
}

// AddStatus adds i to status.
func (auo *AuthorUpdateOne) AddStatus(i int8) *AuthorUpdateOne {
	auo.mutation.AddStatus(i)
	return auo
}

// SetUpdatedAt sets the updated_at field.
func (auo *AuthorUpdateOne) SetUpdatedAt(t time.Time) *AuthorUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetFullName sets the fullName field.
func (auo *AuthorUpdateOne) SetFullName(s string) *AuthorUpdateOne {
	auo.mutation.SetFullName(s)
	return auo
}

// SetAge sets the age field.
func (auo *AuthorUpdateOne) SetAge(i int8) *AuthorUpdateOne {
	auo.mutation.ResetAge()
	auo.mutation.SetAge(i)
	return auo
}

// AddAge adds i to age.
func (auo *AuthorUpdateOne) AddAge(i int8) *AuthorUpdateOne {
	auo.mutation.AddAge(i)
	return auo
}

// AddFilmIDs adds the films edge to Film by ids.
func (auo *AuthorUpdateOne) AddFilmIDs(ids ...uuid.UUID) *AuthorUpdateOne {
	auo.mutation.AddFilmIDs(ids...)
	return auo
}

// AddFilms adds the films edges to Film.
func (auo *AuthorUpdateOne) AddFilms(f ...*Film) *AuthorUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.AddFilmIDs(ids...)
}

// Mutation returns the AuthorMutation object of the builder.
func (auo *AuthorUpdateOne) Mutation() *AuthorMutation {
	return auo.mutation
}

// RemoveFilmIDs removes the films edge to Film by ids.
func (auo *AuthorUpdateOne) RemoveFilmIDs(ids ...uuid.UUID) *AuthorUpdateOne {
	auo.mutation.RemoveFilmIDs(ids...)
	return auo
}

// RemoveFilms removes films edges to Film.
func (auo *AuthorUpdateOne) RemoveFilms(f ...*Film) *AuthorUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.RemoveFilmIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AuthorUpdateOne) Save(ctx context.Context) (*Author, error) {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := author.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}

	var (
		err  error
		node *Author
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthorUpdateOne) SaveX(ctx context.Context) *Author {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AuthorUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthorUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AuthorUpdateOne) sqlSave(ctx context.Context) (a *Author, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   author.Table,
			Columns: author.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: author.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Author.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldStatus,
		})
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldStatus,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: author.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.FullName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: author.FieldFullName,
		})
	}
	if value, ok := auo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldAge,
		})
	}
	if value, ok := auo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: author.FieldAge,
		})
	}
	if nodes := auo.mutation.RemovedFilmsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.FilmsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Author{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{author.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
