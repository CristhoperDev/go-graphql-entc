package cmd

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/CristhoperDev/go-graphql-entc/graphql/generated"
	"github.com/CristhoperDev/go-graphql-entc/graphql/resolver"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "graphql",
	Short: "runs graphql server.",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	log.Fatal(r.Run(fmt.Sprintf(":%s", viper.GetString("APP_PORT"))))
}


// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	config := generated.Config{Resolvers: &resolver.Resolver{}}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

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