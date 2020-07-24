package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/CristhoperDev/go-graphql-entc/graphql/generated"
	"github.com/CristhoperDev/go-graphql-entc/graphql/model"
	"github.com/CristhoperDev/go-graphql-entc/internal/provider/entc"
)

func (r *filmResolver) ID(ctx context.Context, obj *entc.Film) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *filmResolver) Status(ctx context.Context, obj *entc.Film) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *filmResolver) CreatedAt(ctx context.Context, obj *entc.Film) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *filmResolver) UpdatedAt(ctx context.Context, obj *entc.Film) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *filmResolver) Authors(ctx context.Context, obj *entc.Film) ([]*entc.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFilm(ctx context.Context, input model.NewFilm) (*entc.Film, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFilm(ctx context.Context, input model.UpdateFilm) (*entc.Film, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveFilm(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Film(ctx context.Context, id string) (*entc.Film, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllFilm(ctx context.Context, page *int, perPage *int, sortField *string, sortOrder *string, filter *model.FilmFilter) ([]*entc.Film, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllFilmsMeta(ctx context.Context, page *int, perPage *int, filter *model.FilmFilter) (*model.ListMetadata, error) {
	panic(fmt.Errorf("not implemented"))
}

// Film returns generated.FilmResolver implementation.
func (r *Resolver) Film() generated.FilmResolver { return &filmResolver{r} }

type filmResolver struct{ *Resolver }
