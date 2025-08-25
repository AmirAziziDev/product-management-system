CREATE TABLE product_types
(
    id         INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code       INTEGER     NOT NULL UNIQUE CHECK (code >= 0),
    name       TEXT        NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT
ON COLUMN product_types.code IS
  'Stable business code (unsigned int). Used as the first part of SKU.';

CREATE TABLE products
(
    id              INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code            INTEGER     NOT NULL UNIQUE CHECK (code >= 0),
    name            TEXT        NOT NULL UNIQUE,
    description     TEXT,
    product_type_id INTEGER     NOT NULL REFERENCES product_types (id) ON DELETE CASCADE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT
ON COLUMN products.code IS
  'Stable business code (unsigned int). Used as the second part of SKU.';

CREATE TABLE colors
(
    id   INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
    name TEXT    NOT NULL UNIQUE,
    hex  TEXT    NOT NULL CHECK (hex ~ '^#[0-9A-Fa-f]{6}$'),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

COMMENT
ON COLUMN colors.code IS
  'Stable business code (unsigned int). Used as the third part of SKU.';

CREATE TABLE products_colors
(
    product_id INTEGER NOT NULL REFERENCES products (id),
    color_id   INTEGER NOT NULL REFERENCES colors (id),
    PRIMARY KEY (product_id, color_id)
);

CREATE INDEX idx_product_types_code ON product_types (code);
CREATE INDEX idx_product_types_created_at ON product_types (created_at);

CREATE INDEX idx_products_code ON products (code);
CREATE INDEX idx_products_product_type_id ON products (product_type_id);
CREATE INDEX idx_products_created_at ON products (created_at);

CREATE INDEX idx_colors_code ON colors (code);
CREATE INDEX idx_colors_created_at ON colors (created_at);

CREATE INDEX idx_products_colors_product_id ON products_colors (product_id);
CREATE INDEX idx_products_colors_color_id ON products_colors (color_id);
