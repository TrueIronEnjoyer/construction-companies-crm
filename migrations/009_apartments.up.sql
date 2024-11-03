CREATE TABLE IF NOT EXISTS apartments (
    id SERIAL PRIMARY KEY,
    number VARCHAR(50) NOT NULL,
    area DECIMAL(10, 2),
    rooms INT,
    status VARCHAR(100),
    description TEXT
);

CREATE TABLE IF NOT EXISTS apartment_house (
    id SERIAL PRIMARY KEY,
    apartment_id INT REFERENCES apartments(id) ON DELETE CASCADE,
    house_id INT REFERENCES houses(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS apartment_floor (
    id SERIAL PRIMARY KEY,
    apartment_id INT REFERENCES apartments(id) ON DELETE CASCADE,
    floor_id INT REFERENCES floors(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS apartment_entrance (
    id SERIAL PRIMARY KEY,
    apartment_id INT REFERENCES apartments(id) ON DELETE CASCADE,
    entrance_id INT REFERENCES entrances(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS apartment_buyer (
    id SERIAL PRIMARY KEY,
    apartment_id INT REFERENCES apartments(id) ON DELETE CASCADE,
    buyer_id INT REFERENCES buyers(id) ON DELETE CASCADE,
    ownership_share DECIMAL(5, 2) CHECK (ownership_share >= 0 AND ownership_share <= 100)
);

CREATE TABLE IF NOT EXISTS apartment_employee (
    id SERIAL PRIMARY KEY,
    apartment_id INT REFERENCES apartments(id) ON DELETE CASCADE,
    buyer_id INT REFERENCES buyers(id) ON DELETE CASCADE
);