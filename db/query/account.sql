-- name: CreateAccount :one
INSERT INTO accounts (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
