-- +goose Up
CREATE TABLE IF NOT EXISTS orders(
  id BIGSERIAL PRIMARY KEY,
  customer_id BIGSERIAL NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS orders;
