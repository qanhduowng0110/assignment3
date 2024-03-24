package convert

import (
	"assignment3/ent"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"assignment3/graph/model"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.RFC3339, tmpStr)
	}
	return time.Time{}, fmt.Errorf("time should be a string in RFC3339 format")
}

func ConvertIntToTimeStamp(milliseconds int64) time.Time {
	baseTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	// Convert milliseconds to a duration
	duration := time.Duration(milliseconds) * time.Millisecond
	// Add the duration to the base time
	timestamp := baseTime.Add(duration)
	return timestamp

}

func ConvertInt32ToInt(i int32) int {
	return int(i)
}

func convertMapToJSON(obj map[string]interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bytes)
}

func ConvertEntToModel(entEarthquakes []*ent.Earthquake) []*model.Earthquakes {
	modelEarthquakes := make([]*model.Earthquakes, len(entEarthquakes))
	for i, entEarthquake := range entEarthquakes {
		modelEarthquake := &model.Earthquakes{
			ID:         entEarthquake.ID,
			LocationID: entEarthquake.LocationID,
			TimeID:     entEarthquake.TimeID,
			MagitudeID: entEarthquake.MagitudeID,
			URL:        stringPointer(entEarthquake.URL),
			Status:     stringPointer(entEarthquake.Status),
			Tsunami:    stringPointer(strconv.Itoa(int(entEarthquake.Tsunami))),
			Net:        stringPointer(entEarthquake.Net),
			Code:       stringPointer(entEarthquake.Code),
			Sources:    stringPointer(entEarthquake.Sources),
			Nst:        intPointer(entEarthquake.Nst),
			Dmin:       float64Pointer(entEarthquake.Dmin),
			Rms:        float64Pointer(entEarthquake.Rms),
			Gap:        float64Pointer(entEarthquake.Gap),
			CreateAt:   entEarthquake.CreatedAt.Format(time.RFC3339),
			UpdateAt:   entEarthquake.UpdatedAt.Format(time.RFC3339),
		}
		modelEarthquakes[i] = modelEarthquake
	}
	return modelEarthquakes
}

func stringPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func intPointer(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func int32Pointer(i int32) *int32 {
	if i == 0 {
		return nil
	}
	return &i
}

func float64Pointer(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}
