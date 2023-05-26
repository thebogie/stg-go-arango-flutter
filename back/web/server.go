// web/server.go
package web

import (
	"back/internal/usecase"
	"net/http"

	"back/graph"
	"back/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewServer(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Add your authentication middleware here
	//router.Use(middleware.Authentication)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(userUsecase, authUsecase)}))
	router.Handle("/graphql", srv)

	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	return router
}
