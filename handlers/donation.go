package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"

	"demeter/db/generated"

	"net/http"
)

func CreateDonation(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
	ctx.Request().ParseForm()
	for key := range ctx.Request().PostForm {
		if GetValidationErrorMessage(key, ctx.FormValue(key)) != "" {
			// do nothing if any validation errors were found
			// because the validation messages should already be shown in UI
			return ctx.NoContent(406)
		}
	}

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

	locationLat := new(pgtype.Numeric)
	locationLat.Scan("0.0")
	if ctx.FormValue("location-lat") != "" {
		locationLat.Scan(ctx.FormValue("location-lat"))
	}
	locationLong := new(pgtype.Numeric)
	locationLong.Scan("0.0")
	if ctx.FormValue("location-long") != "" {
		locationLong.Scan(ctx.FormValue("location-long"))
	}

	payload := queries.CreateDonationParams{
		Title:           title,
		CreatedByUserID: 1, // TODO:
		StartsAt:        startsAt,
		EndsAt:          endsAt,
		Description:     description,
		Images:          pgtype.Text{}, // TODO:
		ServingsTotal:   total,
		ServingsLeft:    total,
		LocationLat:     *locationLat,
		LocationLong:    *locationLong,
	}
	_, err = query.CreateDonation(dbc, payload)
	if err != nil {
		log.Error(err)
		return ctx.Render(200, "msg-danger", fmt.Sprint("Internal error occurred: ", err.Error()))
	}
	return ctx.Render(200, "msg-success", "Donation post created successfully!")
}

func GetDonationPosts(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
    // get donation post data
    donations, err := query.GetDonations(dbc)
    if err != nil {
        return err
    }

    return ctx.Render(http.StatusOK, "donationPostFeed", donations)
}