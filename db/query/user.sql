-- name: CreateUser :one
INSERT INTO users (email, password)
VALUES ('riw2@example.com', 'pass1234')
RETURNING *;