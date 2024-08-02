package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"github.com/google/uuid"
	"github.com/jamiebmurray25/grpc-crud/database"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "todos.db")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS todos;"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, string(ddl)); err != nil {
		log.Fatal(err)
	}

	todos := []string{
		"Drink a glass of water",
		"Eat a healthy breakfast",
		"Check and respond to emails",
		"Attend scheduled meetings or classes",
		"Complete and submit project report",
		"Review and revise study notes",
		"Read a chapter of a book",
		"Practice a new skill or hobby",
		"Meditate for 10 minutes",
		"Do the laundry",
		"Clean the kitchen",
		"Water the plants",
		"Go for a walk or exercise",
		"Prepare a nutritious lunch",
		"Schedule a medical check-up",
		"Call or message a friend or family member",
		"Plan a weekend outing or activity",
		"Reflect on the day’s accomplishments",
		"Set goals for tomorrow",
		"Unwind with a relaxing activity (e.g., reading, listening to music)",
		"Update personal budget or expenses",
		"Organize digital files or photos",
		"Check and update any ongoing projects or tasks",
	}

	queries := database.New(db)
	for _, todo := range todos {
		uuid, err := uuid.NewV7()

		if err != nil {
			log.Fatal(err)
		}

		_, err = queries.CreateTodo(ctx, database.CreateTodoParams{ID: uuid.String(), Title: todo})

		if err != nil {
			log.Fatal(err)
		}
	}
}
