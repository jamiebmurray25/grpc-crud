package server

import (
	"context"
	"log"
  "errors"

	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

type Server struct {
	pb.UnimplementedTodoServer
	Queries *database.Queries
}

func (s *Server) GetTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodoReply, error) {
	log.Printf("Received: %v", in.GetId())

	todo, err := s.Queries.GetTodo(ctx, in.GetId())

  log.Printf("%+v\n", todo)

	if err != nil {
		return nil, errors.New("Todo not found.")
	}

	return &pb.TodoReply{Title: todo.Title}, nil

}
