-- +goose Up
-- +goose StatementBegin
alter table donations alter column location_lat type numeric(10, 6);
alter table donations alter column location_long type numeric(10, 6);
create index donations_location_idx on donations (location_lat, location_long);

alter table users alter column location_lat type numeric(10, 6);
alter table users alter column location_long type numeric(10, 6);
create index users_location_idx on users (location_lat, location_long);

alter table ratings alter column location_lat type numeric(10, 6);
alter table ratings alter column location_long type numeric(10, 6);
create index ratings_location_idx on users (location_lat, location_long);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
