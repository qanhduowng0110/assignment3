package services

import (
	"assignment3/cmd/api"
	"assignment3/ent"
	"assignment3/middleware"
	"context"

	"github.com/gofiber/fiber/v2"
)

// Get All 100 records
func GetAll(c *fiber.Ctx) []*ent.Earthquake {
	ctx := context.Background()
	Client := api.ConnectDB()
	var earthquakes, errEarth = Client.Earthquake.Query().Limit(100).All(ctx)
	middleware.CheckError(errEarth)
	middleware.SaveRequest(c)
	return earthquakes
}
