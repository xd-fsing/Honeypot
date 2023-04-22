-- name: SaveLoginSpy :one
INSERT INTO spys (kind, ip, request)
VALUES ($1, $2, $3)
RETURNING *;

-- name: TestSql :many
SELECT
FROM spys as s,
     users AS u
where s.created_by = u.id;