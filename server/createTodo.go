package server

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.TodoReply, error) {

	log.Printf("CreateTodo: Attempting to create todo with values: {title: %s}", in.GetTitle())

	todoId, err := uuid.NewV7()

	if err != nil {
		return nil, errors.New("failed to create todo uuid")
	}

	todo, err := s.Queries.CreateTodo(ctx, database.CreateTodoParams{ID: todoId.String(), Title: in.GetTitle()})

	if err != nil {
		return nil, errors.New("failed to create todo")
	}

	log.Printf("CreateTodo: Successfully created todo '%s'", todoId)

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title, Completed: todo.Completed, CreatedAt: timestamppb.New(todo.Createdat)}, nil
}
