CREATE TABLE IF NOT EXISTS stock_transactions (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    sub_variant_id UUID NOT NULL REFERENCES sub_variants(id) ON DELETE CASCADE,
    quantity NUMERIC(20,8) NOT NULL,
    transaction_type VARCHAR(20) NOT NULL,
    transaction_date TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_stock_transactions_product_id ON stock_transactions(product_id);
CREATE INDEX IF NOT EXISTS idx_stock_transactions_sub_variant_id ON stock_transactions(sub_variant_id);
