CREATE TABLE IF NOT EXISTS variants (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    name TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_variants_product_id ON variants(product_id);
