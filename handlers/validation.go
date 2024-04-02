package handlers

import (
	"context"
	"github.com/labstack/echo/v4"

	"demeter/db/generated"
)

func GetValidationErrorMessage(key, val string) string {
	switch key {
	case "title":
		if len(val) < 5 || len(val) > 255 {
			return "Title length has to be between 5 and 255 characters."
		}
	case "description":
		if len(val) < 5 || len(val) > 4096 {
			return "Description length has to be between 5 and 4096 characters."
		}
	}
	return ""
}

func ValidateFormInput(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
	key := ctx.Request().Header.Get("HX-Trigger-Name")
	val := ctx.FormValue(key)
	msg := GetValidationErrorMessage(key, val)
	if msg != "" {
		return ctx.Render(200, "msg-warning", msg)
	}
	return ctx.String(200, "")
}
