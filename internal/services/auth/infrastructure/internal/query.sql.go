// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package internal

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const delete = `-- name: Delete :exec
DELETE FROM
    "refreshToken"
WHERE
    id = $1
`

func (q *Queries) Delete(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, delete, id)
	return err
}

const findByValue = `-- name: FindByValue :one
SELECT
    "createdAt", "updatedAt", id, "userId", value, "clientInfo"
FROM
    "refreshToken"
WHERE
    "value" = $1
`

func (q *Queries) FindByValue(ctx context.Context, value string) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, findByValue, value)
	var i RefreshToken
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.UserId,
		&i.Value,
		&i.ClientInfo,
	)
	return i, err
}

const list = `-- name: List :many
SELECT
    "createdAt", "updatedAt", id, "userId", value, "clientInfo"
FROM
    "refreshToken" LIMIT $1 OFFSET $2
`

type ListParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) List(ctx context.Context, arg ListParams) ([]RefreshToken, error) {
	rows, err := q.db.QueryContext(ctx, list, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RefreshToken
	for rows.Next() {
		var i RefreshToken
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID,
			&i.UserId,
			&i.Value,
			&i.ClientInfo,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const save = `-- name: Save :one
INSERT INTO
    "refreshToken" (
        "createdAt",
        "updatedAt",
        "userId",
        "value",
        "clientInfo"
    )
VALUES
    (NOW(), NOW(), $1, $2, $3) RETURNING "createdAt", "updatedAt", id, "userId", value, "clientInfo"
`

type SaveParams struct {
	UserId     uuid.UUID      `json:"userId"`
	Value      string         `json:"value"`
	ClientInfo sql.NullString `json:"clientInfo"`
}

func (q *Queries) Save(ctx context.Context, arg SaveParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, save, arg.UserId, arg.Value, arg.ClientInfo)
	var i RefreshToken
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID,
		&i.UserId,
		&i.Value,
		&i.ClientInfo,
	)
	return i, err
}