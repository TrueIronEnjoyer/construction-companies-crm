CREATE TABLE IF NOT EXISTS buyers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(50),
    contact_info TEXT,
    description TEXT
);