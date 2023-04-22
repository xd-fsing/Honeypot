-- name: SaveLoginSpy :one
INSERT INTO spys (ip, kind)
VALUES ($1, $2)
RETURNING *;

-- name: TestSql :many
SELECT FROM spys as s,users AS u where s.created_by = u.id;