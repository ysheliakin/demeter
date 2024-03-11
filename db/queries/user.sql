-- name: GetUser :one
select *
from "Users"
where "Id" = @user_id;
