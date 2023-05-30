package repository

import (
	connection "back/internal/db"
	"back/internal/domain"
	"context"
)

type AuthRepository interface {
	FindByEmail(email string) (*domain.User, error)
}

type authRepository struct {
	dbconn *connection.DatabaseConnection
}

func NewAuthRepository(dbconn *connection.DatabaseConnection) AuthRepository {
	return &authRepository{
		dbconn: dbconn}
}

func (r *authRepository) FindByEmail(email string) (*domain.User, error) {
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

	var retuser = &domain.User{}
	_, err = cursor.ReadDocument(ctx, &retuser)
	if err != nil {

	}
	return retuser, err

}
