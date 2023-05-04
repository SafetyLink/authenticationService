-- name: GetUserByID :one
SELECT *
from users
WHERE users.id = $1
LIMIT 1;


-- name: GetUserByEmail :one
SELECT *
from users
WHERE users.email = $1
LIMIT 1;


-- name: GetUserByUsername :one
SELECT *
from users
WHERE users.username = $1
LIMIT 1;

-- name: GetUserFriends :many
SELECT u.id as user_id,
       u.username,
       u.email,
       u.first_name,
       u.last_name,
       u.avatar_id,
       u.created_at,
       u.updated_at
FROM (SELECT friend_id as user1_friend
      FROM friend_list
      WHERE friend_list.user_id = $1
      UNION
      SELECT friend_list.user_id as user1_friend
      FROM friend_list
      WHERE friend_id = $1) as f
         INNER JOIN Users u ON f.user1_friend = u.id
ORDER BY user_id;
