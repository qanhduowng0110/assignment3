package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"assignment3/convert"
	"assignment3/graph"
	"assignment3/graph/model"
	"assignment3/services"
	"context"
)

// ListEarthquakes is the resolver for the listEarthquakes field.
func (r *queryResolver) ListEarthquakes(ctx context.Context, filter *model.EarthquakeFilter, pagination *model.PaginationInput, day *model.DayInput) ([]*model.Earthquakes, error) {
	var limit int
	var offset int
	var mag float64
	var eq_type string
	var place string
	var days int

	if pagination != nil {
		if pagination.Limit != nil {
			limit = *pagination.Limit
		}
		if pagination.Offset != nil {
			offset = *pagination.Offset
		}
	}

	if filter != nil {
		if filter.Magnitude != nil {
			mag = *filter.Magnitude
		}
		if filter.Place != nil {
			place = *filter.Place
		}
		if filter.TypeEarthquake != nil {
			eq_type = *filter.TypeEarthquake
		}
	}

	if day != nil {
		days = *day.Day
	}
	earthquakes_page, errpage := services.QueryEarthquakes(ctx, limit, offset, mag, eq_type, place, days)
	earthquakes_page_convert := convert.ConvertEntToModel(earthquakes_page)
	return earthquakes_page_convert, errpage
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
