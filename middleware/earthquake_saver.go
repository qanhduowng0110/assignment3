package middleware

import (
	"assignment3/cmd/api"
	"assignment3/convert"
	"assignment3/model"
	"context"
	"encoding/json"
	"time"

	fetch "github.com/yngfoxx/gofiber-fetch"
)

var earthquake model.EarthquakeViewModel

func ApiImportEarthquake() {
	Client := api.ConnectDB()
	url := "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_week.geojson"
	res := fetch.Method("GET").FiberFetch(url)
	if res.Error != nil {
		panic(res.Error)
	}
	err := json.Unmarshal([]byte(res.Data.([]byte)), &earthquake)
	CheckError(err)

	ctx := context.Background()
	Client.Earthquake.Delete().Exec(ctx)
	Client.AssociatedEvent.Delete().Exec(ctx)
	Client.Location.Delete().Exec(ctx)
	Client.Magnitude.Delete().Exec(ctx)
	Client.Time.Delete().Exec(ctx)
	Client.Type_eathquake.Delete().Exec(ctx)

	for i := 0; i < len(earthquake.Features); i++ {
		timeNew, err_time := Client.Time.Create().
			SetDateTime(convert.ConvertIntToTimeStamp(earthquake.Features[i].Properties.Time)).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		CheckError(err_time)

		earthquake_time_id := timeNew.ID

		locationNew, err_location := Client.Location.Create().
			SetLongitude(earthquake.Features[i].Geometry.Coordinates[0]).
			SetLatitude(earthquake.Features[i].Geometry.Coordinates[1]).
			SetDept(earthquake.Features[i].Geometry.Coordinates[2]).
			SetPlace(earthquake.Features[i].Properties.Place).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		CheckError(err_location)

		earthquake_location_id := locationNew.ID

		magnitudeNew, err_magnitude := Client.Magnitude.Create().
			SetMagnitudeValue(earthquake.Features[i].Properties.Mag).
			SetMagnitudeType(earthquake.Features[i].Properties.MagType).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		CheckError(err_magnitude)

		earthquake_magnitude_id := magnitudeNew.ID

		earthquakeNew, err_earthquake := Client.Earthquake.Create().
			SetLocationID(earthquake_location_id).
			SetTimeID(earthquake_time_id).
			SetMagnitudeID(earthquake_magnitude_id).
			SetURL(earthquake.Features[i].Properties.URL).
			SetStatus(earthquake.Features[i].Properties.Status).
			SetTsunami(earthquake.Features[i].Properties.Tsunami).
			SetNet(earthquake.Features[i].Properties.Net).
			SetCode(earthquake.Features[i].Properties.Code).
			SetSources(earthquake.Features[i].Properties.Sources).
			SetNst(convert.ConvertInt32ToInt((earthquake.Features[i].Properties.Nst))).
			SetDmin(earthquake.Features[i].Properties.Dmin).
			SetRms(earthquake.Features[i].Properties.Rms).
			SetGap(earthquake.Features[i].Properties.Gap).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		CheckError(err_earthquake)

		earthquake_id := earthquakeNew.ID

		_, err_type := Client.Type_eathquake.Create().
			SetEarthquakeID(earthquake_id).
			SetTypeEathquake(earthquake.Features[i].Properties.Type).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		CheckError(err_type)

		_, err_event := Client.AssociatedEvent.Create().
			SetEarthquakeID(earthquake_id).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		CheckError(err_event)
	}
}
