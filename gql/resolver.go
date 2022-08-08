package gql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/amar-jay/go-api-boilerplate/gql/gen"
)

type Resolver struct{}

// // foo
func (r *mutationResolver) CreateTodo(ctx context.Context, input gen.NewTodo) (*gen.Todo, error) {
	panic("not implemented")
}

// // foo
func (r *queryResolver) Todos(ctx context.Context) ([]*gen.Todo, error) {
	panic("not implemented")
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
