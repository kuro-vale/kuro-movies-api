package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/kuro-vale/kuro-movies-api/graph"
	"github.com/kuro-vale/kuro-movies-api/graph/generated"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	h.Use(extension.FixedComplexityLimit(5))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graph")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
