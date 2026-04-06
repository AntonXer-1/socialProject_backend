CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    photo_url TEXT,
    role VARCHAR(20) NOT NULL CHECK ( role IN ('passenger', 'driver')),
    created_at TIMESTAMP DEFAULT NOW()
);