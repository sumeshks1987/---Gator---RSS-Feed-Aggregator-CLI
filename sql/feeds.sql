-- name: CreateFeed :one
INSERT INTO feeds (
    name,
    url,
    user_id
)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: GetFeedsWithUsers :many
SELECT
    feeds.id,
    feeds.name,
    feeds.url,
    feeds.user_id,
    feeds.created_at,
    feeds.updated_at,
    users.name AS user_name
FROM feeds
JOIN users ON feeds.user_id = users.id
ORDER BY feeds.created_at;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;