// graph/resolver.go
package graph

import (
	"back/graph/generated"
	genmodel "back/graph/generated/model"
	"back/internal/usecase"
	"context"
)

type Resolver struct {
	userUsecase usecase.UserUsecase
	authUsecase usecase.AuthUsecase
}

func NewResolver(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) *Resolver {
	return &Resolver{
		userUsecase: userUsecase,
		authUsecase: authUsecase,
	}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) RegisterUser(ctx context.Context, username string, email string, password string) (*genmodel.AuthPayload, error) {

	// Invoke the corresponding use case method using the generated model types
	result, err := r.Resolver.userUsecase.RegisterUser(username, email, password)
	if err != nil {
		return nil, err
	}

	// Convert the result to the generated AuthPayload type
	authPayload := &genmodel.AuthPayload{
		Token: result.Token,
		User: &genmodel.User{
			ID:       result.User.ID,
			Username: result.User.Username,
			Email:    result.User.Email,
		},
	}

	return authPayload, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, email string, password string) (*genmodel.AuthPayload, error) {

	// Invoke the corresponding use case method using the generated model types
	result, err := r.Resolver.authUsecase.LoginUser(email, password)
	if err != nil {
		return nil, err
	}

	// Convert the result to the generated AuthPayload type
	authPayload := &genmodel.AuthPayload{
		Token: result.Token,
		User: &genmodel.User{
			ID:       result.User.ID,
			Username: result.User.Username,
			Email:    result.User.Email,
		},
	}

	return authPayload, nil

}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (*genmodel.User, error) {
	// Implement the logic to retrieve the authenticated user
	return &genmodel.User{
		ID:       "fish",
		Username: "Mr Toad",
		Email:    "mrtoad@gmail.com",
	}, nil
}
