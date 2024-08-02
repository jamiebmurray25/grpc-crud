package server

import (
	"context"
	"errors"

	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetAllTodos(ctx context.Context, in *pb.Empty) (*pb.GetAllTodosReply, error) {
	todos, err := s.Queries.GetAllTodos(ctx)

	if err != nil {
		return nil, errors.New("failed to fetch todos")
	}

	var todoReplies []*pb.TodoReply

	for _, todo := range todos {
		todoReply := pb.TodoReply{Id: todo.ID, Title: todo.Title, Completed: todo.Completed, CreatedAt: timestamppb.New(todo.Createdat)}

		todoReplies = append(todoReplies, &todoReply)
	}

	return &pb.GetAllTodosReply{Todos: todoReplies}, nil
}
