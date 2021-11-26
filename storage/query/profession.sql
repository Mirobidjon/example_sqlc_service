-- name: CreateProfession :one
INSERT INTO profession(
        id, 
        name,
        description
    )
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetProfession :one
SELECT *
FROM profession
WHERE id = $1;

-- name: GetProfessions :many
SELECT 
    id,
    name,
    description
FROM profession
WHERE name ilike '%' || @search::varchar || '%' 
offset @offset_::integer limit @limit_::integer;

-- name: GetProfessionsCount :one
SELECT count(1)
FROM profession
WHERE name ilike '%' || @search::varchar || '%' ;

-- name: UpdateProfession :exec
UPDATE profession
SET name = $1,
    description = $2
WHERE id = $3;

-- name: DeleteProfession :exec
DELETE
FROM profession
WHERE id = $1;