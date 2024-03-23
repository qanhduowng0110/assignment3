package services

import (
	"assignment3/cmd/api"
	"assignment3/ent"
	"assignment3/ent/location"
	"assignment3/ent/magnitude"
	"assignment3/ent/type_eathquake"
	"assignment3/middleware"
	"assignment3/services"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber"
)

// Filter as required
func QueryEarthquakes(c *fiber.Ctx, pageSize int, pageIndex int, mag float64, eq_type string, place string, daysAgo int) []*ent.Earthquake {
	services.GetAll()
	Client := api.ConnectDB()
	defer api.DB.Close()
	query := Client.Earthquake.Query()
	// Filter by magnitude
	if mag != 0 {
		query.QueryMagnitude().Where(magnitude.MagnitudeValue(mag))
	}
	// Filter by Earthquake type
	if eq_type != "" {
		query.QueryTypes().Where(type_eathquake.TypeEathquake(eq_type))
	}
	// Filter by place
	if place != "" {
		query.QueryLocation().Where(location.Place(place))
	}
	// Filter by days ago
	if daysAgo < 0 {
		panic("daysAgo must be greater than 0")
	} else {
		daynumber := time.Now().AddDate(0, 0, -daysAgo)
		query.QueryTime().Where(sql.FieldGT(time.DateTime, daynumber))
	}
	// pageSize and pageIndex must be greater than 0

	earthquakes_page, errpage := query.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Clone().Order(func(s *sql.Selector) {
		sql.OrderByField("id", sql.OrderDesc())
	}).All(c)
	middleware.CheckError(errpage)
	middleware.SaveRequest(c)
	return earthquakes_page

}
