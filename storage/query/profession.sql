-- name: CreateProfession :one
INSERT INTO profession(
        id, 
        name
    )
VALUES ($1, $2)
RETURNING id;

-- name: GetProfession :one
SELECT *
FROM profession
WHERE id = $1;

-- name: GetProfessions :many
SELECT 
    id,
    name
FROM profession
WHERE name ilike '%' || @search::varchar || '%' 
offset @offset_::integer limit @limit_::integer;

-- name: GetProfessionsCount :one
SELECT count(1)
FROM profession
WHERE name ilike '%' || @search::varchar || '%' ;

-- name: UpdateProfession :exec
UPDATE profession
SET name = $1
WHERE id = $2;

-- name: DeleteProfession :exec
DELETE
FROM profession
WHERE id = $1;