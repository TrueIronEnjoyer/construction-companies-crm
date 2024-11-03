CREATE TABLE IF NOT EXISTS floors (
    id SERIAL PRIMARY KEY,
    number INT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS floor_house (
    id SERIAL PRIMARY KEY,
    floor_iD INT REFERENCES floors(id) ON DELETE CASCADE,
    house_id INT REFERENCES houses(id) ON DELETE CASCADE
);