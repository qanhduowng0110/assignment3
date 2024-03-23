package services

import (
	"assignment3/ent"
	"assignment3/middleware"

	"github.com/gofiber/fiber"
)

// Sync 100 records
func SyncData(c *fiber.Ctx) []*ent.Earthquake {
	middleware.ApiImportEarthquake()
	middleware.SaveRequest(c)
	earthquakes := GetAll()
	middleware.SaveRequest(c)
	return earthquakes
}
