package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"GQLMONGO/graph/model"
	"GQLMONGO/repository"
	"context"
	"math/rand"
	"strconv"
	"GQLMONGO/graph/generated"

	//"github.com/jpastorm/gqlgen-todos/graph/generated"
	//"github.com/jpastorm/gqlgen-todos/graph/model"
	//"github.com/jpastorm/gqlgen-todos/repository"
)

var userRepo repository.UserRepository = repository.New()

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:        strconv.Itoa(rand.Int()),
		Nickname:  input.Nickname,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  input.Password,
	}
	userRepo.Save(user)
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return userRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
