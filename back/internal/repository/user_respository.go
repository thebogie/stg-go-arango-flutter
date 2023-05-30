// internal/repository/user_repository.go
package repository

import (
	connection "back/internal/db"
	"back/internal/domain"
	"context"
)

type UserRepository interface {
	FindByEmail(email string) (*domain.User, error)
	Save(user *domain.User) error
	FindUserByID(id string) (*domain.User, error)
}

type userRepository struct {
	dbconn *connection.DatabaseConnection
}

func NewUserRepository(dbconn *connection.DatabaseConnection) UserRepository {
	return &userRepository{
		dbconn: dbconn}
}

// Implement the UserRepository methods
func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	query := "FOR d IN player FILTER d.email == @email RETURN d"
	bindVars := map[string]interface{}{
		"email": email,
	}
	//TODO: pass context from top level?
	ctx := context.Background()
	cursor, err := r.dbconn.Db.Query(ctx, query, bindVars)
	if err != nil {

	}
	defer cursor.Close()
	return &domain.User{}, nil

	// Implement the logic to find a user by email
	return nil, nil
}

func (r *userRepository) FindUserByID(userid string) (*domain.User, error) {
	query := "FOR d IN player FILTER d.email == @email RETURN d"
	bindVars := map[string]interface{}{
		"email": userid,
	}
	//TODO: pass context from top level?
	ctx := context.Background()
	cursor, err := r.dbconn.Db.Query(ctx, query, bindVars)
	if err != nil {

	}
	defer cursor.Close()
	return &domain.User{}, nil

	// Implement the logic to find a user by email
	return nil, nil
}

func (r *userRepository) Save(user *domain.User) error {
	// Implement the logic to save a user
	return nil
}
