-- name: CreateUser :one
INSERT INTO users (name)
VALUES ($1)
RETURNING *;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;

-- name: GetAnyUser :one
SELECT *
FROM users
LIMIT 1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;