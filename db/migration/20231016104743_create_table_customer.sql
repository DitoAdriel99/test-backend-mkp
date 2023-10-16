-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers
(
    id              uuid        not null primary key,
    customer_name   varchar        not null,
    contact_info    varchar     not null,
    created_at      timestamp   not null,
    updated_at      timestamp
);
INSERT INTO customers (id, customer_name, contact_info, created_at, updated_at)
VALUES
  ('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'John Doe', 'john.doe@example.com', NOW(), NOW()),
  ('44b5f9a3-2e66-4379-8f54-8d5c19536a2a', 'Jane Smith', 'jane.smith@example.com', NOW(), NOW()),
  ('a242231e-fd14-4de0-82c1-f80a0a5d689b', 'Michael Johnson', 'michael.johnson@example.com', NOW(), NOW()),
  ('e7464567-9e8f-45a9-89a7-2c89ac1e0d3c', 'Anna Davis', 'anna.davis@example.com', NOW(), NOW()),
  ('12d98a3a-7827-4e54-9be5-7aef1542b7e2', 'David Wilson', 'david.wilson@example.com', NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customers
-- +goose StatementEnd
