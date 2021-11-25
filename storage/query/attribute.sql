-- name: CreateAttribute :one
INSERT INTO attribute (
        id, 
        name, 
        type
    )
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetAttribute :one
SELECT *
FROM attribute
WHERE id = $1;

-- name: GetAttributes :many
SELECT *
FROM attribute
WHERE name ilike '%' || @search::varchar || '%'
offset @offset_ limit @limit_;

-- name: GetAttributesCount :one
SELECT count(1)
FROM attribute
WHERE name ilike '%' || @search::varchar || '%';

-- name: UpdateAttribute :exec
UPDATE attribute
SET name = $1,
    type = $2
WHERE id = $3;

-- name: DeleteAttribute :exec
DELETE
FROM attribute
WHERE id = $1;
