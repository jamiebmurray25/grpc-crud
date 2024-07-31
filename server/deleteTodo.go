package server

import (
	"context"
	"errors"

	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

func (s *Server) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.Empty, error) {
	_, err := s.Queries.GetTodoById(ctx, in.Id)

	if err != nil {
		return nil, errors.New("todo does not exist")
	}

	err = s.Queries.DeleteTodo(ctx, in.Id)

	if err != nil {
		return nil, errors.New("failed to delete todo")
	}

	return &pb.Empty{}, nil
}
