CREATE TABLE IF NOT EXISTS sub_variants (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    option_ids TEXT[] NOT NULL,
    sku TEXT NOT NULL UNIQUE,
    stock NUMERIC(20,8) DEFAULT 0
);
CREATE INDEX IF NOT EXISTS idx_sub_variants_product_id ON sub_variants(product_id);
