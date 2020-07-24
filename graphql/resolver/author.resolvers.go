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

func (r *authorResolver) ID(ctx context.Context, obj *entc.Author) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *authorResolver) Status(ctx context.Context, obj *entc.Author) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *authorResolver) CreatedAt(ctx context.Context, obj *entc.Author) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *authorResolver) UpdatedAt(ctx context.Context, obj *entc.Author) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *authorResolver) Age(ctx context.Context, obj *entc.Author) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *authorResolver) Films(ctx context.Context, obj *entc.Author) ([]*entc.Film, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*entc.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, input model.UpdateAuthor) (*entc.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveAuthor(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Author(ctx context.Context, id string) (*entc.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllAuthor(ctx context.Context, page *int, perPage *int, sortField *string, sortOrder *string, filter *model.AuthorFilter) ([]*entc.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllAuthorMeta(ctx context.Context, page *int, perPage *int, filter *model.AuthorFilter) (*model.ListMetadata, error) {
	panic(fmt.Errorf("not implemented"))
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
