package middleware

import (
	"assignment3/cmd/api"
	"assignment3/convert"
	"assignment3/model"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

var request_report model.RequestViewModel
var m = make(map[string]interface{})

func SaveRequest(c *fiber.Ctx) {
	Client := api.ConnectDB()
	ctx := context.Background()
	for key, value := range c.GetReqHeaders() {
		m[key] = value
	}

	_, err := Client.Request.Create().
		SetRequestURL(c.OriginalURL()).
		SetRequestMethod(c.Method()).
		SetRequestHeaders(m).
		SetRequestBody(convert.ConvertByteToMap(c.Request().Body())).
		SetResponseStatusCode(int32(c.Response().StatusCode())).
		SetResponseBody(convert.ConvertByteToMap(c.Response().Body())).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	CheckError(err)
}
