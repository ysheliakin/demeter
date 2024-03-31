-- name: CreateDonation :one
insert into donations(
    title, created_by_user_id, created_at, starts_at, ends_at, description, images, servings_total, servings_left, location_lat, location_long
)
values (
    $1, NOW(), $2, $3, $4, $5, $6, $7, $8, $9, $10
)
returning id;
