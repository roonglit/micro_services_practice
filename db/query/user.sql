-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ('riw@example.com', '382e0360e4eb7b70034fbaa69bec5786'),
       ('yao3@example.com', '382e0360e4eb7b70034fbaa69bec5786')
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

