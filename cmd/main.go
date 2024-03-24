package main

import (
	"assignment3/cmd/api"
	"assignment3/router"
	"os"

	_ "entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const defaultPort = "3000"

var App *fiber.App

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	api.ConnectDB()
	defer api.DB.Close()

	// Check authentication

	// Set Routes
	App := fiber.New()
	router.SetRoutes(App)
	App.Listen(":3000")

}
