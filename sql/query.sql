-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: GetTodoById :one
SELECT * FROM todos WHERE id = ? LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (id, title) VALUES (?, ?) RETURNING *;

-- name: UpdateTodo :one
UPDATE todos SET title = ?, completed = ? WHERE id = ? RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?; 