CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(40) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- INSERT INTO users (email, password)
-- VALUES ('riw@example.com', '382e0360e4eb7b70034fbaa69bec5786'),
--        ('yao3@example.com', '382e0360e4eb7b70034fbaa69bec5786');