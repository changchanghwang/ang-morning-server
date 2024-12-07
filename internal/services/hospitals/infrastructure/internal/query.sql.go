// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package internal

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const countByCity = `-- name: CountByCity :one
SELECT COUNT(*) FROM hospital 
WHERE city = ANY($1::text[])
`

func (q *Queries) CountByCity(ctx context.Context, dollar_1 []string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countByCity, pq.Array(dollar_1))
	var count int64
	err := row.Scan(&count)
	return count, err
}

const delete = `-- name: Delete :exec
UPDATE "hospital" SET "deletedAt" = NOW()
WHERE id = $1
`

func (q *Queries) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, delete, id)
	return err
}

const findByCity = `-- name: FindByCity :many
SELECT "createdAt", "updatedAt", "deletedAt", id, name, phone, city, "roadAddress", latitude, longitude, "zipCode" FROM hospital 
WHERE city = ANY($1::text[])
`

func (q *Queries) FindByCity(ctx context.Context, dollar_1 []string) ([]Hospital, error) {
	rows, err := q.db.QueryContext(ctx, findByCity, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hospital
	for rows.Next() {
		var i Hospital
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID,
			&i.Name,
			&i.Phone,
			&i.City,
			&i.RoadAddress,
			&i.Latitude,
			&i.Longitude,
			&i.ZipCode,
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

const list = `-- name: List :many
SELECT "createdAt", "updatedAt", "deletedAt", id, name, phone, city, "roadAddress", latitude, longitude, "zipCode" FROM "hospital" WHERE "deletedAt" IS NULL
`

func (q *Queries) List(ctx context.Context) ([]Hospital, error) {
	rows, err := q.db.QueryContext(ctx, list)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hospital
	for rows.Next() {
		var i Hospital
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID,
			&i.Name,
			&i.Phone,
			&i.City,
			&i.RoadAddress,
			&i.Latitude,
			&i.Longitude,
			&i.ZipCode,
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
INSERT INTO "hospital" (
  "createdAt", "updatedAt", id, name, phone, city, "roadAddress", latitude, longitude, "zipCode"
) VALUES (
  NOW(), NOW(), $1, $2, $3, $4, $5, $6, $7, $8
) ON CONFLICT ("id") DO UPDATE
SET 
  "updatedAt" = NOW(),
  name = EXCLUDED.name, 
  phone = EXCLUDED.phone, 
  city = EXCLUDED.city,
  "roadAddress" = EXCLUDED."roadAddress",
  latitude = EXCLUDED.latitude,
  longitude = EXCLUDED.longitude,
  "zipCode" = EXCLUDED."zipCode"
RETURNING "createdAt", "updatedAt", "deletedAt", id, name, phone, city, "roadAddress", latitude, longitude, "zipCode"
`

type SaveParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	City        string    `json:"city"`
	RoadAddress string    `json:"roadAddress"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	ZipCode     string    `json:"zipCode"`
}

func (q *Queries) Save(ctx context.Context, arg SaveParams) (Hospital, error) {
	row := q.db.QueryRowContext(ctx, save,
		arg.ID,
		arg.Name,
		arg.Phone,
		arg.City,
		arg.RoadAddress,
		arg.Latitude,
		arg.Longitude,
		arg.ZipCode,
	)
	var i Hospital
	err := row.Scan(
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ID,
		&i.Name,
		&i.Phone,
		&i.City,
		&i.RoadAddress,
		&i.Latitude,
		&i.Longitude,
		&i.ZipCode,
	)
	return i, err
}