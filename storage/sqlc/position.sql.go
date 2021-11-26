// Code generated by sqlc. DO NOT EDIT.
// source: position.sql

package sqlc

import (
	"context"
)

const createPosition = `-- name: CreatePosition :one
INSERT INTO position (
        id, 
        name, 
        profession_id
    )
VALUES ($1, $2, $3)
RETURNING id
`

type CreatePositionParams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ProfessionID string `json:"profession_id"`
}

func (q *Queries) CreatePosition(ctx context.Context, arg CreatePositionParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createPosition, arg.ID, arg.Name, arg.ProfessionID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deletePosition = `-- name: DeletePosition :exec
DELETE
FROM position
WHERE id = $1
`

func (q *Queries) DeletePosition(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePosition, id)
	return err
}

const getPosition = `-- name: GetPosition :one
SELECT id, name, profession_id
FROM position
WHERE id = $1
`

func (q *Queries) GetPosition(ctx context.Context, id string) (Position, error) {
	row := q.db.QueryRowContext(ctx, getPosition, id)
	var i Position
	err := row.Scan(&i.ID, &i.Name, &i.ProfessionID)
	return i, err
}

const getPositions = `-- name: GetPositions :many
SELECT id, name, profession_id
FROM position
WHERE name ilike '%' || $1::varchar || '%' AND 
CASE  WHEN $2::varchar = '' THEN true ELSE profession_id = $2::varchar END
offset $3 limit $4
`

type GetPositionsParams struct {
	Search       string `json:"search"`
	ProfessionID string `json:"profession_id"`
	Offset       int32  `json:"offset_"`
	Limit        int32  `json:"limit_"`
}

func (q *Queries) GetPositions(ctx context.Context, arg GetPositionsParams) ([]Position, error) {
	rows, err := q.db.QueryContext(ctx, getPositions,
		arg.Search,
		arg.ProfessionID,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Position
	for rows.Next() {
		var i Position
		if err := rows.Scan(&i.ID, &i.Name, &i.ProfessionID); err != nil {
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

const getPositionsCount = `-- name: GetPositionsCount :one
SELECT count(1)
FROM position
WHERE name ilike '%' || $1::varchar || '%' AND 
CASE  WHEN $2::varchar = '' THEN true ELSE profession_id = $2::varchar END
`

type GetPositionsCountParams struct {
	Search       string `json:"search"`
	ProfessionID string `json:"profession_id"`
}

func (q *Queries) GetPositionsCount(ctx context.Context, arg GetPositionsCountParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPositionsCount, arg.Search, arg.ProfessionID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updatePosition = `-- name: UpdatePosition :exec
UPDATE position
SET name          = $1,
    profession_id = $2
WHERE id = $3
`

type UpdatePositionParams struct {
	Name         string `json:"name"`
	ProfessionID string `json:"profession_id"`
	ID           string `json:"id"`
}

func (q *Queries) UpdatePosition(ctx context.Context, arg UpdatePositionParams) error {
	_, err := q.db.ExecContext(ctx, updatePosition, arg.Name, arg.ProfessionID, arg.ID)
	return err
}
