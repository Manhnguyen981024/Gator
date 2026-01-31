-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES(
        $1,$2,$3,$4,$5
    )
    RETURNING *
)
SELECT feed_follows.*
, users.name AS user_name
, feeds.name AS feed_name
, feeds.url  AS feed_url
FROM inserted_feed_follow AS feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE feeds.url = $1;

-- name: GetFeedFollowsForUsers :many
SELECT feeds.name as feed_name, users.name as User_name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollowByURL :exec
DELETE FROM feed_follows
WHERE user_id = $1 and feed_id=$2;