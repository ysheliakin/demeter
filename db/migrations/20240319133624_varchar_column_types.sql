-- +goose Up
-- +goose StatementBegin
alter table users alter column email type varchar(255);
alter table users alter column avatar_url type varchar(1024);
alter table users alter column images type varchar(4096);
alter table users alter column auth_provider type varchar(32);
alter table users alter column auth_hash type varchar(1024);
alter table users alter column about type varchar(4096);

alter table donations add column title varchar(255) not null default 'Local event';
alter table donations alter column description type varchar(4096);
alter table donations alter column images type varchar(4096);
alter table donations alter column description type varchar(4096);

alter table donation_events alter column comment type varchar(1024);
alter table donation_events alter column images type varchar(4096);

alter table ratings alter column comment type varchar(1024);
alter table ratings alter column images type varchar(4096);

alter table notifications alter column comment type varchar(1024);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
