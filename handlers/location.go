package handlers

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"

	"demeter/db/generated"
)

func f64(val pgtype.Numeric) float64 {
	ret, _ := val.Float64Value()
	return ret.Float64
}

func num(val float64) pgtype.Numeric {
	ret := new(pgtype.Numeric)
	ret.Scan(strconv.FormatFloat(val, 'f', -1, 64))
	return *ret
}

func nums(val string) pgtype.Numeric {
	ret := new(pgtype.Numeric)
	ret.Scan(val)
	return *ret
}

// TODO: add unit tests for these
func getCoordsRange(lat, long pgtype.Numeric, milesRange float64) (minLat, maxLat, minLong, maxLong pgtype.Numeric) {
	milesEarthRadius := 3958.8
	degs2rads := math.Pi / 180.0
	rads2degs := 180.0 / math.Pi

	_lat := f64(lat)
	latDelta := milesRange / milesEarthRadius * rads2degs
	minLat, maxLat = num(_lat-latDelta), num(_lat+latDelta)

	_long := f64(long)
	longDelta := milesRange / (milesEarthRadius * math.Cos(_lat*degs2rads)) * rads2degs
	minLong, maxLong = num(_long-longDelta), num(_long+longDelta)

	return
}

type LocationMessage struct {
	Message string `json:"message"`
}

func GetLocationResults(dbc context.Context, query *queries.Queries, ctx echo.Context, log echo.Logger) error {
	lat, long := ctx.QueryParam("lat"), ctx.QueryParam("long")
	if lat == "" || long == "" {
		return ctx.JSON(400, LocationMessage{"Latitude or longitude value is empty."})
	}

	kind := ctx.QueryParam("kind")
	minLat, maxLat, minLong, maxLong := getCoordsRange(nums(lat), nums(long), 30.0)

	switch kind {

	case "users":
		params := queries.GetUsersInRangeParams{MinLat: minLat, MaxLat: maxLat, MinLong: minLong, MaxLong: maxLong}
		results, err := query.GetUsersInRange(dbc, params)
		if err != nil {
			return ctx.JSON(422, LocationMessage{fmt.Sprintf("Unable to get users in requested range: %s", err.Error())})
		}
		return ctx.JSON(200, results)

	case "donations":
		params := queries.GetDonationsInRangeParams{MinLat: minLat, MaxLat: maxLat, MinLong: minLong, MaxLong: maxLong}
		results, err := query.GetDonationsInRange(dbc, params)
		if err != nil {
			return ctx.JSON(422, LocationMessage{fmt.Sprintf("Unable to get donations in requested range: %s", err.Error())})
		}
		return ctx.JSON(200, results)

	default:
		return ctx.JSON(400, LocationMessage{"Impossible kind provided in query params."})
	}
}
