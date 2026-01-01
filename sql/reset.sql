-- name: Reset :exec
TRUNCATE TABLE posts, feeds, users
RESTART IDENTITY
CASCADE;