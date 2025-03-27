CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);
INSERT INTO persons (email, phone, first_name, last_name) VALUES
    ('john.doe@example.com', '1234567890', 'John', 'Doe'),
    ('jane.doe@example.com', '0987654321', 'Jane', 'Doe'),
    ('alice.smith@example.com', '1112223333', 'Alice', 'Smith'),
    ('bob.jones@example.com', '4445556666', 'Bob', 'Jones'),
    ('charlie.brown@example.com', '7778889999', 'Charlie', 'Brown'),
    ('david.wilson@example.com', '1231231234', 'David', 'Wilson'),
    ('eva.miller@example.com', '5675675678', 'Eva', 'Miller'),
    ('frank.moore@example.com', '8908908901', 'Frank', 'Moore'),
    ('grace.taylor@example.com', '2342342345', 'Grace', 'Taylor'),
    ('harry.anderson@example.com', '6786786789', 'Harry', 'Anderson');