package server

import (
	"context"
	"errors"

	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

func (s *Server) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.TodoReply, error) {
	todo, err := s.Queries.UpdateTodo(ctx, database.UpdateTodoParams{Title: in.Title, ID: in.Id})

	if err != nil {
		return nil, errors.New("failed to update todo")
	}

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title}, nil
}
