CREATE TABLE IF NOT EXISTS order_items (
    order_item_id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    estimate_id uuid NOT NULL REFERENCES estimates(estimate_id),
    price int NOT NULL CHECK (price >= 1),
    quantity int NOT NULL CHECK (quantity >= 1),
    amount int NOT NULL CHECK (amount >= 1),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);