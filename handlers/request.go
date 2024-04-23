package handlers

import (
	"context"
	queries "demeter/db/generated"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func RequestForm(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	post, err := query.GetDonation(dbc, int32(id))
	if err != nil {
		return err
	}
	return ctx.Render(200, "request", post)
}

func CreateRequest(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	comment := ctx.FormValue("comment")
	payload := queries.CreateRequestParams{
		DonationID:  pgtype.Int4{Int32: int32(id)},
		RequesterID: pgtype.Int4{Int32: 1}, // TODO: add users
		Comment:     pgtype.Text{String: comment},
	}
	fmt.Println(payload)
	_, err = query.CreateRequest(dbc, payload)
	if err != nil {
		log.Error(err)
		return ctx.Render(200, "msg-danger", fmt.Sprint("Internal error occurred: ", err.Error()))
	}
	return ctx.Render(200, "msg-success", "Request created successfully!")
}
