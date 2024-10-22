CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(15) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO users (email, password)
VALUES ('riw@example.com', 'pass1234'),
       ('yao3@example.com', 'pass1234');