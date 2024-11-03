CREATE TABLE IF NOT EXISTS houses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    year_built INT,
    description TEXT
);

CREATE TABLE IF NOT EXISTS house_block (
    id SERIAL PRIMARY KEY,
    block_id INT REFERENCES blocks(id) ON DELETE CASCADE,
    house_id INT REFERENCES houses(id) ON DELETE CASCADE
);