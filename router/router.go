package router

import (
	"assignment3/middleware"
	"assignment3/services"
	"context"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/sync", func(c *fiber.Ctx) error {
		earthquakes := services.SyncData(c)
		c.SendString("Sucess!")
		return c.JSON(earthquakes)
	})

	app.Get("/get-all", func(c *fiber.Ctx) error {
		earthquakes := services.GetAll(c)
		return c.JSON(earthquakes)
	})

	app.Get("/filter", func(c *fiber.Ctx) error {
		ctx := context.Background()
		earthquakes := services.GetAll(c)
		pageSize, errSize := strconv.Atoi(c.Query("pageSize"))
		middleware.CheckError(errSize)
		pageIndex, errIndex := strconv.Atoi(c.Query("pageIndex"))
		middleware.CheckError(errIndex)
		mag, errMag := strconv.ParseFloat(c.Query("mag"), 64)
		middleware.CheckError(errMag)
		eqType := c.Query("eqType")
		place := c.Query("place")
		daysAgo, errDaysAgo := strconv.Atoi(c.Query("daysAgo"))
		middleware.CheckError(errDaysAgo)
		earthquakes, _ = services.QueryEarthquakes(ctx, pageIndex, pageSize, mag, eqType, place, daysAgo)
		c.SendString("Sucess!")
		return c.JSON(earthquakes)
	})
}
