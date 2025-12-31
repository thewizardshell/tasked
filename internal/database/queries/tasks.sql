-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;

-- name: ListTasksByUser :many
SELECT * FROM tasks
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateTask :one
INSERT INTO tasks (title, description, status, priority, user_id, due_date)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, description = $3, status = $4, priority = $5, due_date = $6, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateTaskStatus :one
UPDATE tasks
SET status = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;
