DO $$
BEGIN
  IF (SELECT COUNT(*) FROM products) = 0 THEN
    COPY products (code, name, description) FROM '/seeds/products.txt' 
    WITH (FORMAT csv, DELIMITER ';', NULL '');
    RAISE NOTICE 'Database initialized with % products', (SELECT COUNT(*) FROM products);
  ELSE
    RAISE NOTICE 'Database already contains % products, skipping initial seed', (SELECT COUNT(*) FROM products);
  END IF;
END $$;