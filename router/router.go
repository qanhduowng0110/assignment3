package router

import (
	"assignment3/middleware"
	"assignment3/services"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	app.Get("/sync", func(c *fiber.Ctx) error {
		earthquakes := services.SyncData()
		c.SendString("Sucess!")
		return c.JSON(earthquakes)
	})

	app.Get("/get-all", func(c *fiber.Ctx) error {
		earthquakes := services.GetAll(c)
		return c.JSON(earthquakes)
	})

	app.Get("/filter", func(c *fiber.Ctx) error {
		earthquakes := services.GetAll(c)
		pageSize, errSize := strconv.Atoi(c.Query("pageSize"))
		middleware.CheckError(errSize)
		pageIndex, errIndex := strconv.Atoi(c.Query("pageIndex"))
		middleware.CheckError(errIndex)
		mag, errMag := strconv.ParseFloat(c.Query("mag"), 64)
		middleware.CheckError(errMag)
		eqType, errEqtype := c.Query("eqType")
		middleware.CheckError(errEqtype)
		place, errPlace := c.Query("place")
		middleware.CheckError(errPlace)
		daysAgo, errDaysAgo := strconv.Atoi(c.Query("daysAgo"))
		middleware.CheckError(errDaysAgo)
		earthquakes = services.QueryEarthquakes(c, pageIndex, pageSize, mag, eqType, place, daysAgo)
		c.SendString("Sucess!")
		return c.JSON(earthquakes)
	})
}
