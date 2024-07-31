package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"github.com/jamiebmurray25/grpc-crud/server"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func setupDatabase() (*database.Queries, error) {
	db, err := sql.Open("sqlite3", "todos.db")

	if err != nil {
		return nil, err
	}

	queries := database.New(db)

	return queries, nil
}

func main() {
	fmt.Printf("Starting server on port :3333\n")

	queries, err := setupDatabase()

	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":3333")

	if err != nil {
		log.Fatal(err)
	}

	registrar := grpc.NewServer()

	pb.RegisterTodoServer(registrar, &server.Server{Queries: queries})

	reflection.Register(registrar)

	if err := registrar.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
