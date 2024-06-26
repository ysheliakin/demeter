// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: location.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getDonationsInRange = `-- name: GetDonationsInRange :many
select id, created_by_user_id, created_at, starts_at, ends_at, description, images, servings_total, servings_left, location_lat, location_long, title
from donations
where location_lat between $1 and $2
    and location_long between $3 and $4
order by ((location_lat - ($2+@min_lat)/2) + (location_long - ($3+@max_long)/2))
`

type GetDonationsInRangeParams struct {
	MinLat  pgtype.Numeric
	MaxLat  pgtype.Numeric
	MinLong pgtype.Numeric
	MaxLong pgtype.Numeric
}

func (q *Queries) GetDonationsInRange(ctx context.Context, arg GetDonationsInRangeParams) ([]Donation, error) {
	rows, err := q.db.Query(ctx, getDonationsInRange,
		arg.MinLat,
		arg.MaxLat,
		arg.MinLong,
		arg.MaxLong,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Donation
	for rows.Next() {
		var i Donation
		if err := rows.Scan(
			&i.ID,
			&i.CreatedByUserID,
			&i.CreatedAt,
			&i.StartsAt,
			&i.EndsAt,
			&i.Description,
			&i.Images,
			&i.ServingsTotal,
			&i.ServingsLeft,
			&i.LocationLat,
			&i.LocationLong,
			&i.Title,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersInRange = `-- name: GetUsersInRange :many
select id, is_organization, name, email, about, avatar_url, images, auth_provider, auth_hash, registered_at, last_visited_at, is_online, location_lat, location_long
from users
where location_lat between $1 and $2
    and location_long between $3 and $4
order by ((location_lat - ($2+@min_lat)/2) + (location_long - ($3+@max_long)/2))
`

type GetUsersInRangeParams struct {
	MinLat  pgtype.Numeric
	MaxLat  pgtype.Numeric
	MinLong pgtype.Numeric
	MaxLong pgtype.Numeric
}

func (q *Queries) GetUsersInRange(ctx context.Context, arg GetUsersInRangeParams) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsersInRange,
		arg.MinLat,
		arg.MaxLat,
		arg.MinLong,
		arg.MaxLong,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.IsOrganization,
			&i.Name,
			&i.Email,
			&i.About,
			&i.AvatarUrl,
			&i.Images,
			&i.AuthProvider,
			&i.AuthHash,
			&i.RegisteredAt,
			&i.LastVisitedAt,
			&i.IsOnline,
			&i.LocationLat,
			&i.LocationLong,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
