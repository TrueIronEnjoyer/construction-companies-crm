CREATE TABLE blocks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS block_complex (
    id SERIAL PRIMARY KEY,
    block_id INT REFERENCES blocks(id), ON DELETE CASCADE
    complex_id INT REFERENCES complexes(id) ON DELETE CASCADE
);