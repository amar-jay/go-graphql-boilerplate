package gql

import (
	"context"
	// "errors"
	"github.com/amar-jay/go-api-boilerplate/gql/gen"
)

// // foo
func (r *mutationResolver) Register(ctx context.Context, input gen.RegisterLogin) (*gen.RegisterLoginOutput, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) Login(ctx context.Context, input gen.RegisterLogin) (*gen.RegisterLoginOutput, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) UpdateUser(ctx context.Context, input gen.UpdateUser) (*gen.User, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
	panic("not implemented")
}

// // foo
func (r *mutationResolver) ResetPassword(ctx context.Context, token string, password string) (*gen.RegisterLoginOutput, error) {
	panic("not implemented")
}
