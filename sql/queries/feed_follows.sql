-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
from inserted_feed_follow
INNER JOIN users on inserted_feed_follow.user_id = users.id
INNER JOIN feeds on inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowForUser :many
SELECT feeds.name, users.name
FROM feed_follows
INNER JOIN users on feed_follows.user_id = users.id
INNER JOIN feeds on feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds, users
WHERE feed_follows.user_id = users.id AND
feed_follows.feed_id = feeds.id AND
feed_follows.user_id = $1 AND
feeds.url = $2;
