-- name: CreateGoal :one
INSERT INTO goals (
  user_id, name, description, type, target_value
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetGoal :one
SELECT * FROM goals
WHERE id = $1 LIMIT 1;

-- name: ListGoals :many
SELECT * FROM goals
WHERE user_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateGoal :one
UPDATE goals
SET name = $2, description = $3, type = $4, target_value = $5, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteGoal :exec
DELETE FROM goals
WHERE id = $1;
