-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts p
INNER JOIN feeds f ON p.feed_id = f.id
INNER JOIN users u on f.user_id = u.id
WHERE u.id = $1
ORDER BY p.published_at DESC
LIMIT $2;

-- name: GetPostByUrl :one
SELECT * FROM posts
WHERE url = $1;
