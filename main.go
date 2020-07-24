package main

import (
	"context"
	"github.com/CristhoperDev/go-graphql-entc/graphql/generated"
	"github.com/CristhoperDev/go-graphql-entc/graphql/resolver"
	"github.com/CristhoperDev/go-graphql-entc/internal/database"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	client, err := database.Open()
	if err != nil {
		panic(err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	log.Fatal(r.Run(":3000"))
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
