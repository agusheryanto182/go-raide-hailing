CREATE TABLE IF NOT EXISTS merchants (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar(30) NOT NULL,
    merchant_category varchar(30) NOT NULL,
    image_url varchar NOT NULL,
    location_lat FLOAT8 NOT NULL,
    location_long FLOAT8 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS merchant_items (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    merchant_id uuid REFERENCES merchants(id),
    name VARCHAR(30) NOT NULL,
    product_category VARCHAR(30) NOT NULL,
    price int NOT NULL,
    image_url VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DO $$ 
BEGIN
    CREATE EXTENSION IF NOT EXISTS postgis;
    CREATE EXTENSION IF NOT EXISTS postgis_topology;
EXCEPTION
    WHEN others THEN
        RAISE NOTICE 'PostGIS is not available';
END $$;