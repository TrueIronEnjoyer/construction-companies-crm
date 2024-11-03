CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(50),
    contact_info TEXT,
    description TEXT
);

CREATE TABLE IF NOT EXISTS employee_developer (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,
    developer_id INT REFERENCES developers(id) ON DELETE CASCADE,
    position VARCHAR(100),
    start_date DATE
);