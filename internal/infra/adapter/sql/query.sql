-- name: GetUserByID :one
SELECT username,
       email,
       first_name,
       last_name,
       avatar_id,
       created_at,
       updated_at
from users
WHERE users.id = $1
LIMIT 1;

-- name: GetUserSecurityByEmail :one
SELECT users.id,
       users.username,
       users.email,
       users.first_name,
       users.last_name,
       users.avatar_id,
       users.created_at,
       users.updated_at,
       account_security.password,
       account_security.updated_at as account_security_updated_at,
       account_security.device_id  as account_device_id
from users
         INNER JOIN account_security ON users.id = account_security.user_id
WHERE users.email = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
from users
WHERE users.username = $1
LIMIT 1;

-- name: GetSelf :many
SELECT u.id,
       u.username,
       u.email,
       u.first_name,
       u.last_name,
       u.avatar_id,
       u.created_at,
       u.updated_at,
       chat.chat_id,
       chat.unread_message,
       chat.last_message_at,
       chat.viewed,
       chat.viewed_at,
       friend.id        AS friend_id,
       friend.username  AS friend_username,
       friend.avatar_id AS friend_avatar_id
FROM users u
         INNER JOIN chat ON (u.id = chat.user_id OR u.id = chat.friend_id)
         INNER JOIN users friend ON (chat.user_id = friend.id OR chat.friend_id = friend.id) AND friend.id != u.id
WHERE u.id = $1;


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
      FROM friend
      WHERE friend.user_id = $1
      UNION
      SELECT friend.user_id as user1_friend
      FROM friend
      WHERE friend_id = $1) as f
         INNER JOIN Users u ON f.user1_friend = u.id
ORDER BY user_id;
