package server

import (
	"context"
	"log"
  "errors"

	"github.com/jamiebmurray25/grpc-crud/database"
  "github.com/google/uuid"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

type Server struct {
	pb.UnimplementedTodoServer
	Queries *database.Queries
}

func (s *Server) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.TodoReply, error) {
	log.Printf("Received: %v", in.GetId())

	todo, err := s.Queries.GetTodo(ctx, in.GetId())

  log.Printf("%+v\n", todo)

	if err != nil {
		return nil, errors.New("Todo not found.")
	}

  return &pb.TodoReply{Id: todo.ID, Title: todo.Title}, nil
}


func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.TodoReply, error) {

  todoId, err := uuid.NewV7()
  
  if err != nil {
    return nil, errors.New("Failed to create todo")
  }

	log.Printf("Created todo with uuid: %s", todoId.String())

  todo, err := s.Queries.CreateTodo(ctx, database.CreateTodoParams{ID: todoId.String(), Title: in.Title})

  if err != nil {
    log.Printf(err.Error());
    return nil, errors.New("Failed to create todo")
  }

  return &pb.TodoReply{Id: todo.ID, Title: todo.Title}, nil
}
