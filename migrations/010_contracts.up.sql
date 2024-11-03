CREATE TABLE IF NOT EXISTS contracts (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    cost DECIMAL(10, 2),
    status VARCHAR(100),
    description TEXT
);


CREATE TABLE IF NOT EXISTS contract_buyer (
    id SERIAL PRIMARY KEY,
    contract_id INT REFERENCES contracts(id) ON DELETE CASCADE,
    buyer_id INT REFERENCES buyers(id) ON DELETE CASCADE, 
    ownership_share DECIMAL(5, 2) CHECK (ownership_share >= 0 AND ownership_share <= 100)
);


CREATE TABLE IF NOT EXISTS contract_employee (
    id SERIAL PRIMARY KEY,
    contractid INT REFERENCES contracts(id) ON DELETE CASCADE,
    employeeid INT REFERENCES employees(id) ON DELETE CASCADE
);