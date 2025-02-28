-- name: CreateUser :one
INSERT INTO users (
  account_id, email, full_name
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE account_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET full_name = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
