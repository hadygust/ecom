-- +goose Up
CREATE TABLE IF NOT EXISTS order_items(
  id BIGSERIAL PRIMARY KEY,
  order_id BIGSERIAL NOT NULL,
  product_id BIGSERIAL NOT NULL,
  quantity INTEGER NOT NULL CHECK (quantity > 0),
  price_in_cent INTEGER NOT NULL CHECK(price_in_cent >= 0),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS order_items;
