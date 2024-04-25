-- name: GetUsersInRange :many
select *
from users
where location_lat between @min_lat and @max_lat
    and location_long between @min_long and @max_long
order by ((location_lat - (@max_lat+@min_lat)/2) + (location_long - (@min_long+@max_long)/2));

-- name: GetDonationsInRange :many
select *
from donations
where location_lat between @min_lat and @max_lat
    and location_long between @min_long and @max_long
order by ((location_lat - (@max_lat+@min_lat)/2) + (location_long - (@min_long+@max_long)/2));
