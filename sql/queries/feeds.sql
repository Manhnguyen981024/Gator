-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, url, users.name as user_name
FROM feeds
    JOIN users ON feeds.user_id = users.id;

-- name: MarkfeedFetched :exec
UPDATE feeds
SET updated_at = $1
, last_fetched_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT id
, created_at
, updated_at
, name
, url
, user_id
, last_fetched_at
FROM feeds
WHERE last_fetched_at IS NULL
AND user_id = $1
ORDER BY last_fetched_at  DESC
FETCH FIRST 1 ROW ONLY;