CREATE TABLE complexes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS complex_deloper (
    id SERIAL PRIMARY KEY,
    complex_id INT REFERENCES complexes(id) ON DELETE CASCADE,
    developer_id INT REFERENCES developers(id) ON DELETE CASCADE
);