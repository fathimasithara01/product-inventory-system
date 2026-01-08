CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    product_id BIGINT NOT NULL UNIQUE,
    product_code TEXT NOT NULL UNIQUE,
    product_name TEXT NOT NULL,
    product_image TEXT,
    created_date TIMESTAMPTZ DEFAULT now(),
    updated_date TIMESTAMPTZ DEFAULT now(),
    created_user UUID NOT NULL,
    is_favourite BOOLEAN DEFAULT false,
    active BOOLEAN DEFAULT true,
    hsn_code TEXT,
    total_stock NUMERIC(20,8) DEFAULT 0
);
