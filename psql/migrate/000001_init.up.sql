CREATE TABLE contacts (
    id VARCHAR(200) NOT NULL UNIQUE,
    name VARCHAR(200) NOT NULL,
    username VARCHAR(200) NOT NULL UNIQUE,
    age VARCHAR(10) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    profession VARCHAR(100) NOT NULL,
    is_admin BOOLEAN DEFAULT false
);
