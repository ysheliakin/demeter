-- name: GetUser :one
select *
from users
where id = @user_id;
