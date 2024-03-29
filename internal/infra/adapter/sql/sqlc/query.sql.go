// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getSelf = `-- name: GetSelf :many
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
         LEFT JOIN chat ON (u.id = chat.user_id OR u.id = chat.friend_id)
         LEFT JOIN users friend ON (chat.user_id = friend.id OR chat.friend_id = friend.id) AND friend.id != u.id
WHERE u.id = $1
`

type GetSelfRow struct {
	ID             int64
	Username       string
	Email          string
	FirstName      pgtype.Text
	LastName       pgtype.Text
	AvatarID       pgtype.Int8
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
	ChatID         pgtype.Int8
	UnreadMessage  pgtype.Int8
	LastMessageAt  pgtype.Timestamp
	Viewed         pgtype.Bool
	ViewedAt       pgtype.Timestamp
	FriendID       pgtype.Int8
	FriendUsername pgtype.Text
	FriendAvatarID pgtype.Int8
}

func (q *Queries) GetSelf(ctx context.Context, id int64) ([]GetSelfRow, error) {
	rows, err := q.db.Query(ctx, getSelf, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSelfRow
	for rows.Next() {
		var i GetSelfRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.AvatarID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ChatID,
			&i.UnreadMessage,
			&i.LastMessageAt,
			&i.Viewed,
			&i.ViewedAt,
			&i.FriendID,
			&i.FriendUsername,
			&i.FriendAvatarID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByID = `-- name: GetUserByID :one
SELECT username,
       email,
       first_name,
       last_name,
       avatar_id,
       created_at,
       updated_at
from users
WHERE users.id = $1
LIMIT 1
`

type GetUserByIDRow struct {
	Username  string
	Email     string
	FirstName pgtype.Text
	LastName  pgtype.Text
	AvatarID  pgtype.Int8
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) GetUserByID(ctx context.Context, id int64) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, first_name, last_name, avatar_id, created_at, updated_at
from users
WHERE users.username = $1
LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserFriends = `-- name: GetUserFriends :many
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
ORDER BY user_id
`

type GetUserFriendsRow struct {
	UserID    int64
	Username  string
	Email     string
	FirstName pgtype.Text
	LastName  pgtype.Text
	AvatarID  pgtype.Int8
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) GetUserFriends(ctx context.Context, userID int64) ([]GetUserFriendsRow, error) {
	rows, err := q.db.Query(ctx, getUserFriends, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserFriendsRow
	for rows.Next() {
		var i GetUserFriendsRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.FirstName,
			&i.LastName,
			&i.AvatarID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserSecurityByEmail = `-- name: GetUserSecurityByEmail :one
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
LIMIT 1
`

type GetUserSecurityByEmailRow struct {
	ID                       int64
	Username                 string
	Email                    string
	FirstName                pgtype.Text
	LastName                 pgtype.Text
	AvatarID                 pgtype.Int8
	CreatedAt                pgtype.Timestamp
	UpdatedAt                pgtype.Timestamp
	Password                 string
	AccountSecurityUpdatedAt pgtype.Timestamp
	AccountDeviceID          int64
}

func (q *Queries) GetUserSecurityByEmail(ctx context.Context, email string) (GetUserSecurityByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserSecurityByEmail, email)
	var i GetUserSecurityByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.AvatarID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Password,
		&i.AccountSecurityUpdatedAt,
		&i.AccountDeviceID,
	)
	return i, err
}
