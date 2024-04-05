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
  deleted_at INTEGER DEFAULT 0
);

