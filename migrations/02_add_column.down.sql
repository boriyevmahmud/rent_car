ALTER TABLE customers
DROP CONSTRAINT customers_deleted_at_phone_unique;

ALTER TABLE cars
DROP COLUMN price;

ALTER TABLE customers
DROP COLUMN password;

ALTER TABLE customers
DROP CONSTRAINT unique_login;

ALTER TABLE customers
DROP COLUMN login;
