// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package queries

import (
	"context"
)

const getUser = `-- name: GetUser :one
select id, is_organization, name, email, about, avatar_url, images, auth_provider, auth_hash, registered_at, last_visited_at, is_online, location_lat, location_long
from users
where id = $1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
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
	)
	return i, err
}
