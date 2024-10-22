-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ('riw2@example.com', 'pass1234')
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = 'riw2@example.com' LIMIT 1;

