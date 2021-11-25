-- name: CreatePositionAttribute :one
INSERT INTO position_attributes (
        id, 
        attribute_id, 
        position_id, 
        value
    )
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetPositionAttribute :one
SELECT *
FROM position_attributes
WHERE id = $1;

-- name: GetPositionAttributes :many
SELECT 
    pa.id,
    pa.attribute_id,
    pa.position_id,
    pa.value,
    a.name as attribute_name,
    a.type as attribute_type
FROM position_attributes pa
JOIN attribute a ON pa.attribute_id = a.id
WHERE position_id = $1 ;

-- name: UpdatePositionAttribute :exec
UPDATE position_attributes
SET 
    attribute_id = $1,
    position_id  = $2,
    value        = $3
WHERE id = $4;

-- name: DeletePositionAttribute :exec
DELETE
FROM position_attributes
WHERE position_id = $1;