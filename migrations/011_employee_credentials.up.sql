CREATE TABLE IF NOT EXISTS employee_credentials (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL ON DELETE CASCADE,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);

CREATE TABLE IF NOT EXISTS employe_object_access (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id) ON DELETE CASCADE,
    object_type VARCHAR NOT NULL,
    read_access BOOLEAN,
    update_access BOOLEAN,
    delete_access BOOLEAN
)