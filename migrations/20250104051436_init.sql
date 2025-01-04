-- +goose Up
-- +goose StatementBegin
create table if not exists public.persons (
    id serial primary key,
    last_name varchar(255) not null,
    first_name varchar(255) not null,
    middle_name varchar(255),
    age integer not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.persons;
-- +goose StatementEnd
