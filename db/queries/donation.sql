-- name: CreateDonation :one
insert into donations(
    title, created_by_user_id, created_at, starts_at, ends_at, description, images, servings_total, servings_left, location_lat, location_long
)
values (
    $1, $2, NOW(), $3, $4, $5, $6, $7, $8, $9, $10
)
returning id;

-- name: GetDonations :many
SELECT * 
FROM donations
ORDER BY id DESC;

-- name: GetDonation :one
select *
from donations
where id = @id;

-- name: CreateRequest :one
insert into donation_requests(donation_id, requester_id, comment, created_at)
values (@donation_id, @requester_id, @comment, NOW())
returning id;

