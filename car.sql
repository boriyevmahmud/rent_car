CREATE TABLE IF NOT EXISTS cars (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name Varchar(50) NOT NULL,
    brand Varchar(20) NOT NULL,
    model Varchar(30) NOT NULL,
    hourse_power INTEGER DEFAULT 0,
    colour VARCHAR(20) NOT NULL DEFAULT 'black',
    engine_cap DECIMAL(10,2) NOT NULL DEFAULT 1.0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

