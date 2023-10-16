-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products 
(
    id              uuid        not null primary key,
    product_name    varchar     not null,
    product_price   int,
    stock           int,
    created_at      timestamp   not null,
    updated_at      timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products
-- +goose StatementEnd
