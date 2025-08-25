DO $$
DECLARE
  line TEXT;
  parts TEXT[];
BEGIN
  -- Seed product types first
  IF (SELECT COUNT(*) FROM product_types) = 0 THEN
    -- Read from product_types.txt
    FOR line IN SELECT unnest(string_to_array(
      pg_read_file('/seeds/product_types.txt'), E'\n'
    )) LOOP
      CONTINUE WHEN trim(line) = '';
      parts := string_to_array(trim(line), ';');
      INSERT INTO product_types (code, name) VALUES (parts[1]::INTEGER, parts[2]);
    END LOOP;
    RAISE NOTICE 'Database initialized with % product types', (SELECT COUNT(*) FROM product_types);
  ELSE
    RAISE NOTICE 'Database already contains % product types, skipping seed', (SELECT COUNT(*) FROM product_types);
  END IF;

  -- Seed products
  IF (SELECT COUNT(*) FROM products) = 0 THEN
    -- Read from products.txt
    FOR line IN SELECT unnest(string_to_array(
      pg_read_file('/seeds/products.txt'), E'\n'
    )) LOOP
      CONTINUE WHEN trim(line) = '';
      parts := string_to_array(trim(line), ';');
      INSERT INTO products (code, name, description, product_type_id) VALUES 
        (parts[1]::INTEGER, parts[2], NULLIF(parts[3], ''), parts[4]::INTEGER);
    END LOOP;
    RAISE NOTICE 'Database initialized with % products', (SELECT COUNT(*) FROM products);
  ELSE
    RAISE NOTICE 'Database already contains % products, skipping initial seed', (SELECT COUNT(*) FROM products);
  END IF;

  -- Seed colors
  IF (SELECT COUNT(*) FROM colors) = 0 THEN
    -- Read from colors.txt
    FOR line IN SELECT unnest(string_to_array(
      pg_read_file('/seeds/colors.txt'), E'\n'
    )) LOOP
      CONTINUE WHEN trim(line) = '';
      parts := string_to_array(trim(line), ';');
      INSERT INTO colors (code, name, hex) VALUES (parts[1]::INTEGER, parts[2], parts[3]);
    END LOOP;
    RAISE NOTICE 'Database initialized with % colors', (SELECT COUNT(*) FROM colors);
  ELSE
    RAISE NOTICE 'Database already contains % colors, skipping color seed', (SELECT COUNT(*) FROM colors);
  END IF;

  -- Seed product-color associations
  IF (SELECT COUNT(*) FROM products_colors) = 0 THEN
    -- Read from products_colors.txt
    FOR line IN SELECT unnest(string_to_array(
      pg_read_file('/seeds/products_colors.txt'), E'\n'
    )) LOOP
      CONTINUE WHEN trim(line) = '';
      parts := string_to_array(trim(line), ';');
      INSERT INTO products_colors (product_id, color_id) VALUES (parts[1]::INTEGER, parts[2]::INTEGER);
    END LOOP;
    RAISE NOTICE 'Database initialized with % product-color associations', (SELECT COUNT(*) FROM products_colors);
  ELSE
    RAISE NOTICE 'Database already contains % product-color associations, skipping association seed', (SELECT COUNT(*) FROM products_colors);
  END IF;
END $$;