-- +goose Up
-- +goose StatementBegin
create table if not exists public.payments (
    id serial primary key,
    person_id integer references public.persons (id),
    sum numeric(10, 2),
    created timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists public.payments;
-- +goose StatementEnd
