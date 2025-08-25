CREATE TABLE product_types (
    id SERIAL PRIMARY KEY,
    code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
    name TEXT UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now()
);

COMMENT ON COLUMN product_types.code IS
  'Stable business code (unsigned int). Used as the first part of SKU.';

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
    name TEXT NOT NULL UNIQUE,
    description TEXT NULL,
    product_type_id INTEGER REFERENCES product_types(id),
    created_at TIMESTAMPTZ DEFAULT now()
);

COMMENT ON COLUMN products.code IS
  'Stable business code (unsigned int). Used as the second part of SKU.';

CREATE INDEX idx_product_types_code ON product_types(code);
CREATE INDEX idx_products_code ON products(code);
CREATE INDEX idx_products_product_type_id ON products(product_type_id);
CREATE INDEX idx_products_created_at ON products(created_at DESC);