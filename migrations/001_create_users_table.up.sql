CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'verified', 'suspend');
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    status user_status NOT NULL DEFAULT 'inactive',
    tel TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP 
);