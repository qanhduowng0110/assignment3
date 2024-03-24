package main

import (
	"assignment3/cmd/api"
	"assignment3/graph"
	"assignment3/middleware"
	"assignment3/resolver"
	"assignment3/router"
	"net/http"
	"os"

	_ "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	_ "github.com/lib/pq"
)

const defaultPort = "3000"

var App *fiber.App
var C *fiber.Ctx

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	api.ConnectDB()
	defer api.DB.Close()

	App := fiber.New()
	// Check authentication
	App.Use(middleware.MiddlewareAuthenticateJWT())
	middleware.SaveRequest(C)
	// Set Routes
	router.SetRoutes(App)

	// GraphQL schema resolver handler.
	resolverHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
	// Handler for GraphQL Playground
	playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")

	App.Post("/graphql", adaptor.HTTPHandler(resolverHandler))
	App.Options("/graphql", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})
	// Enable playground for query/testing
	App.Get("/", adaptor.HTTPHandler(playgroundHandler))
	App.Listen(":3000")

}
