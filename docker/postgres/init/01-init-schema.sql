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

CREATE TABLE colors (
    id SERIAL PRIMARY KEY,
    code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
    name TEXT UNIQUE,
    hex CHAR(7) NOT NULL CHECK (hex ~ '^#[0-9A-Fa-f]{6}$')
);

COMMENT ON COLUMN colors.code IS
  'Stable business code (unsigned int). Used as the third part of SKU.';

CREATE TABLE products_colors (
    product_id INTEGER REFERENCES products(id),
    color_id INTEGER REFERENCES colors(id),
    PRIMARY KEY(product_id, color_id)
);

CREATE INDEX idx_product_types_code ON product_types(code);
CREATE INDEX idx_products_code ON products(code);
CREATE INDEX idx_products_product_type_id ON products(product_type_id);
CREATE INDEX idx_products_created_at ON products(created_at DESC);
CREATE INDEX idx_colors_code ON colors(code);
CREATE INDEX idx_products_colors_product_id ON products_colors(product_id);
CREATE INDEX idx_products_colors_color_id ON products_colors(color_id);