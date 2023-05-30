package web

import (
	"back/internal/usecase"
	"back/web/middleware/auth"
	"github.com/go-chi/chi/middleware"
	"net/http"

	"back/graph"
	"back/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func NewServer(userUsecase usecase.UserUsecase, authUsecase usecase.AuthUsecase) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	router.Use(auth.Middleware(userUsecase))
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(userUsecase, authUsecase)}))

	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	return router
}
