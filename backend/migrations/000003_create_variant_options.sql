CREATE TABLE IF NOT EXISTS variant_options (
    id UUID PRIMARY KEY,
    variant_id UUID NOT NULL REFERENCES variants(id) ON DELETE CASCADE,
    value TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_variant_options_variant_id ON variant_options(variant_id);
