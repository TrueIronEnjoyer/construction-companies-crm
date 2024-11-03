CREATE TABLE IF NOT EXISTS entrances (
    id SERIAL PRIMARY KEY,
    number INT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS entrance_house (
    id SERIAL PRIMARY KEY,
    entrance_id INT REFERENCES entrances(id) ON DELETE CASCADE,
    house_id INT REFERENCES houses(id) ON DELETE CASCADE
);