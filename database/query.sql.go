// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package database

import (
	"context"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (id, title) VALUES (?, ?) RETURNING id, title, completed, createdat
`

type CreateTodoParams struct {
	ID    string
	Title string
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.ID, arg.Title)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.Createdat,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getAllTodos = `-- name: GetAllTodos :many
SELECT id, title, completed, createdat FROM todos
`

func (q *Queries) GetAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Completed,
			&i.Createdat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodoById = `-- name: GetTodoById :one
SELECT id, title, completed, createdat FROM todos WHERE id = ? LIMIT 1
`

func (q *Queries) GetTodoById(ctx context.Context, id string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.Createdat,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos SET title = ?, completed = ? WHERE id = ? RETURNING id, title, completed, createdat
`

type UpdateTodoParams struct {
	Title     string
	Completed bool
	ID        string
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.Title, arg.Completed, arg.ID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.Createdat,
	)
	return i, err
}
