package server

import (
	"context"
	"errors"
	"log"

	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

func (s *Server) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.Empty, error) {

	log.Printf("DeleteTodo: Attempting to delete todo with uuid: '%s'", in.GetId())

	_, err := s.Queries.GetTodoById(ctx, in.GetId())

	if err != nil {
		return nil, errors.New("todo does not exist")
	}

	err = s.Queries.DeleteTodo(ctx, in.GetId())

	if err != nil {
		return nil, errors.New("failed to delete todo")
	}

	log.Printf("DeleteTodo: Successfully deleted todo '%s'", in.GetId())

	return &pb.Empty{}, nil
}
