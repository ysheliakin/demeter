package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"

	"demeter/db/generated"
)

// TODO: these probably need to be moved to a different place, not in donation.go
func getValidationErrorMessage(key, val string) string {
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
	msg := getValidationErrorMessage(key, val)
	if msg != "" {
		return ctx.Render(200, "msg-warning", msg)
	}
	return ctx.String(200, "")
}

func CreateDonation(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
	// TODO: use getValidationErrorMessage for all fields to check.
	//	     If all good, then submit
	//		 Otherwise, return only error messages (btw, how?)
	validationErrors := make(map[string]string)

	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	startsAt := pgtype.Timestamp{Valid: false}
	_startsAt, err := time.Parse(time.RFC3339, ctx.FormValue("starts-at"))
	if err == nil {
		startsAt = pgtype.Timestamp{Time: _startsAt}
	}
	endsAt := pgtype.Timestamp{Valid: false}
	_endsAt, err := time.Parse(time.RFC3339, ctx.FormValue("ends-at"))
	if err == nil {
		endsAt = pgtype.Timestamp{Time: _endsAt}
	}
	total := pgtype.Int4{Valid: false}
	if _total, err := strconv.Atoi(ctx.FormValue("servings-total")); err == nil {
		total = pgtype.Int4{Int32: int32(_total)}
	}

	payload := queries.CreateDonationParams{
		Title:         title,
		StartsAt:      startsAt,
		EndsAt:        endsAt,
		Description:   description,
		Images:        pgtype.Text{}, // TODO:
		ServingsTotal: total,
		ServingsLeft:  total,
		LocationLat:   0.0, // TODO:
		LocationLong:  0.0, // TODO:
	}
	_, err = query.CreateDonation(dbc, payload)
	if err != nil {
		log.Errorf("failed to created a donation: %s, error: %s\n")
	} else if len(validationErrors) == 0 {
		return ctx.Render(201, "msg-success", "Successfully created new donation posting!")
	}

	return ctx.Render(200, "donate", nil)
}
