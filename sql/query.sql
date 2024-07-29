-- name: GetTodo :one
SELECT * FROM todos WHERE id = ? LIMIT 1;
