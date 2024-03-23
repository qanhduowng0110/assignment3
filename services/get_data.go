package services

import (
	"assignment3/cmd/api"
	"assignment3/ent"
	"assignment3/middleware"

	"github.com/gofiber/fiber"
)

// Get All 100 records
func GetAll(c *fiber.Ctx) []*ent.Earthquake {
	Client := api.ConnectDB()
	defer api.DB.Close()
	var earthquakes, errEarth = Client.Earthquake.Query().Limit(100).All(c)
	middleware.CheckError(errEarth)
	middleware.SaveRequest(c)
	return earthquakes
}
