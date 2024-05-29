CREATE TABLE IF NOT EXISTS estimates(
    estimate_id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(user_id),
    user_lat FLOAT8 not null,
    user_lon FLOAT8 not null,
    total_price int not null,
    estimated_delivery_time int not null,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS estimates_estimate_id ON estimates (estimate_id);    