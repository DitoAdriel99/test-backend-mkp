-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sales
(
    id          uuid        not null primary key,
    customer_id uuid        not null,
    product_id  uuid        not null,
    quantity    int        not null,
    created_at  timestamp   not null,
    updated_at  timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sales
-- +goose StatementEnd
