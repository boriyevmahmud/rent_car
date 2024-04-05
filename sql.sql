CREATE TABLE IF NOT EXISTS cars (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name Varchar(50) NOT NULL,
  year INTEGER NOT NULL,
  brand Varchar(20) NOT NULL,
  model Varchar(30) NOT NULL,
  horse_power INTEGER DEFAULT 0,
  colour VARCHAR(20) NOT NULL DEFAULT 'black',
  engine_cap DECIMAL(10,2) NOT NULL DEFAULT 1.0,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS customers (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50),
  email VARCHAR(50) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  login VARCHAR(255) UNIQUE,
  password VARCHAR(255) NOT NULL,
  address VARCHAR(20) NOT NULL,
  is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at INTEGER DEFAULT 0,
  CONSTRAINT customers_deleted_at_phone_unique UNIQUE (deleted_at, phone)
);

CREATE TABLE orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  car_id UUID REFERENCES cars(id),
  customer_id UUID REFERENCES customers(id),
  from_date DATE NOT NULL,
  to_date DATE NOT NULL,
  status VARCHAR(255) NOT NULL,
  payment_status BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at INTEGER DEFAULT 0,
);

-- Unique constraint for deleted-at and phone
ALTER TABLE customers
ADD CONSTRAINT customers_deleted_at_phone_unique UNIQUE (deleted_at, phone);

ALTER TABLE cars
ADD COLUMN price DECIMAL(10, 2);

ALTER TABLE orders
ADD COLUMN deleted_at INTEGER DEFAULT 0;

ALTER TABLE cars
RENAME COLUMN hourse_power TO horse_power;

ALTER TABLE customers
ADD COLUMN password TYPE VARCHAR(255);

ALTER TABLE customers
ALTER COLUMN password SET NOT NULL;

ALTER TABLE customers
ADD COLUMN login VARCHAR(255);

ALTER TABLE customers
ALTER COLUMN login SET NOT NULL;

ALTER TABLE customers
ADD CONSTRAINT unique_login UNIQUE (login);

INSERT INTO cars (name, brand, year, model, hourse_power, colour, engine_cap)
VALUES ('Honda', 'Brandy', 2023, 'Model1', 150, 'Red', 1.8);

INSERT INTO cars (name, brand, year, model, hourse_power, colour, engine_cap)
VALUES ('Kia', 'CCH', 2020, 'Model2', 200, 'Blue', 2.0);

INSERT INTO customers (first_name, last_name, email, phone, address)
VALUES ('Sheila', 'Tursunova', 'sht@gmail.com', '+998776655444', 'Tashkent24');

INSERT INTO customers (first_name, last_name, email, phone, address)
VALUES ('Surayyo', 'Sanoyeva', 'sanoyeva@gmail.com', '998776665522', 'Sirdaryo');

SELECT * FROM cars;
SELECT * FROM customers;
SELECT * FROM orders;

-- Inserting data into cars table
INSERT INTO cars (name, year, brand, model, hourse_power, colour, engine_cap) VALUES 
('Car 1', 2020, 'Brand 1', 'Model 1', 120, 'Red', 1.6),
('Car 2', 2021, 'Brand 2', 'Model 2', 130, 'Blue', 1.8),
('Car 3', 2023, 'Brand 30', 'Model 30', 200, 'Black', 2.0);

-- Inserting data into customers table
INSERT INTO customers (first_name, last_name, email, phone, address) VALUES 
('Customer 1', 'Lastname 1', 'customer1@gmail.com', '9984567890', 'Address 1'),
('Customer 2', 'Lastname 2', 'customer2@gmail.com', '0987654321', 'Address 2'),
('Customer 3', 'Lastname 30', 'customer30@gmail.com', '1122334455', 'Address 30');

INSERT INTO customers (first_name, last_name, email, phone, address) VALUES 
('Customer 4', 'Lastname 4', 'customer4@gmail.com', '9984567800', 'Address 2673'),
('Customer 5', 'Lastname 5', 'customer5@gmail.com', '0987654301', 'Address 34'),
('Customer 6', 'Lastname 6', 'customer60@gmail.com', '9981334405', 'Address 56'),
('Customer 7', 'Lastname 7', 'customer7@gmail.com', '9984567844', 'Address 788'),
('Customer 8', 'Lastname 8', 'customer8@gmail.com', '0987654300', 'Address 20'),
('Customer 9', 'Lastname 9', 'customer90@gmail.com', '9982334415', 'Address 380'),
('Customer 10', 'Lastname 10', 'customer10@gmail.com', '9984565890', 'Address 4572'),
('Customer 11', 'Lastname 11', 'customer11@gmail.com', '9987654322', 'Address 46'),
('Customer 12', 'Lastname 12', 'customer12@gmail.com', '9982334450', 'Address 370');

-- Inserting data into orders table
INSERT INTO orders (car_id, customer_id, from_date, to_date, status, payment_status) VALUES 
('f60253d2-cb9a-486a-a6d7-e384363587ab', '91a548c9-4b55-4ed2-96c4-3ce7c61284a8', '2024-01-01', '2024-01-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '91a548c9-4b55-4ed2-96c4-3ce7c61284a8', '2024-02-01', '2024-02-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '91a548c9-4b55-4ed2-96c4-3ce7c61284a8', '2024-12-01', '2024-12-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '91a548c9-4b55-4ed2-96c4-3ce7c61284a8', '2024-02-01', '2024-02-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '91a548c9-4b55-4ed2-96c4-3ce7c61284a8', '2024-02-01', '2024-02-10', 'active', TRUE);

INSERT INTO orders (car_id, customer_id, from_date, to_date, status, payment_status) VALUES 
('f60253d2-cb9a-486a-a6d7-e384363587ab', '1bd4b190-80e3-47c7-8cf6-255a3567cd75', '2024-01-01', '2024-01-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '1bd4b190-80e3-47c7-8cf6-255a3567cd75', '2024-02-01', '2024-02-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '1bd4b190-80e3-47c7-8cf6-255a3567cd75', '2024-12-01', '2024-12-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '1bd4b190-80e3-47c7-8cf6-255a3567cd75', '2024-02-01', '2024-02-10', 'active', TRUE),
('f60253d2-cb9a-486a-a6d7-e384363587ab', '1bd4b190-80e3-47c7-8cf6-255a3567cd75', '2024-02-01', '2024-02-10', 'active', TRUE);

SELECT COUNT(id) FROM customers;
SELECT COUNT(id) FROM orders;

--SELECT COUNT(DISTINCT o.car_id) FROM orders o WHERE o.customer_id = $1;
SELECT COUNT(distinct(o.car_id)) FROM customers c, orders o WHERE o.customer_id = c.id;

--SELECT COUNT(o.car_id) FROM orders o WHERE o.customer_id = $1;
SELECT COUNT(o.car_id) FROM orders o WHERE o.customer_id = 'id';

