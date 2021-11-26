-- name: CreatePosition :one
INSERT INTO position (
        id, 
        name, 
        profession_id
    )
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetPosition :one
SELECT *
FROM position
WHERE id = $1;

-- name: GetPositions :many
SELECT *
FROM position
WHERE name ilike '%' || @search::varchar || '%' AND 
CASE  WHEN @profession_id::varchar = '' THEN true ELSE profession_id = @profession_id::varchar END
offset @offset_ limit @limit_;

-- name: GetPositionsCount :one
SELECT count(1)
FROM position
WHERE name ilike '%' || @search::varchar || '%' AND 
CASE  WHEN @profession_id::varchar = '' THEN true ELSE profession_id = @profession_id::varchar END
;

-- name: UpdatePosition :exec
UPDATE position
SET name          = $1,
    profession_id = $2
WHERE id = $3;

-- name: DeletePosition :exec
DELETE
FROM position
WHERE id = $1;