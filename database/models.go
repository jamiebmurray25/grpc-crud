// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"
)

type Todo struct {
	ID        string
	Title     string
	Completed bool
	Createdat time.Time
}
