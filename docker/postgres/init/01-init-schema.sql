-- Database initialization script
-- This file is executed automatically when the PostgreSQL container starts

-- Create products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    code INTEGER NOT NULL UNIQUE CHECK (code >= 0),
    name TEXT NOT NULL UNIQUE,
    description TEXT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Add comment explaining the code column usage for SKU
COMMENT ON COLUMN products.code IS
  'Stable business code (unsigned int). Used as the second part of SKU (model_code).';

-- Create indexes for better performance
CREATE INDEX idx_products_code ON products(code);
CREATE INDEX idx_products_created_at ON products(created_at DESC);