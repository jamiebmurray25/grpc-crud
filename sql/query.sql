-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: GetTodoById :one
SELECT * FROM todos WHERE id = ? LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (id, title) VALUES (?, ?) RETURNING *;
