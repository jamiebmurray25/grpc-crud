package server

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.TodoReply, error) {

	todoId, err := uuid.NewV7()

	if err != nil {
		return nil, errors.New("failed to create todo uuid")
	}

	log.Printf("Creating todo with uuid: %s", todoId.String())

	todo, err := s.Queries.CreateTodo(ctx, database.CreateTodoParams{ID: todoId.String(), Title: in.Title})

	if err != nil {
		return nil, errors.New("failed to create todo")
	}

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title}, nil
}
