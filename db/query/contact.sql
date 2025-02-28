-- name: CreateContact :one
INSERT INTO contacts (
  user_id, first_name, last_name, email, phone, company_name, address
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetContact :one
SELECT * FROM contacts
WHERE id = $1 LIMIT 1;

-- name: ListContacts :many
SELECT * FROM contacts
WHERE user_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateContact :one
UPDATE contacts
SET first_name = $2, last_name = $3, email = $4, phone = $5, company_name = $6, address = $7, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM contacts
WHERE id = $1;
