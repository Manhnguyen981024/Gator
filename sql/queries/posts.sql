-- name: CreatePost :one
INSERT INTO posts(created_at, title, url, description, published_at, feed_id)
	VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPostsByUserId :many
SELECT posts.title
, posts.url
, posts.description
, feeds.id 
, feeds.user_id
, posts.published_at
, feeds."name" as feed_name
FROM posts
JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN users on feeds.user_id = users.id
WHERE users.id = $1
ORDER BY published_at DESC
LIMIT $2;
