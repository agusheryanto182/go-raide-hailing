CREATE TABLE IF NOT EXISTS order_items (
    order_item_id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    estimate_id uuid NOT NULL REFERENCES estimates(estimate_id),
    order_id uuid NULL REFERENCES orders(order_id),
    merchant_id uuid NOT NULL REFERENCES merchants(merchant_id),
    merchant_item_id uuid NOT NULL REFERENCES merchant_items(item_id),
    quantity int NOT NULL CHECK (quantity >= 1),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);