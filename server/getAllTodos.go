package server

import (
	"context"
	"errors"
	"log"

	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetAllTodos(ctx context.Context, in *pb.Empty) (*pb.GetAllTodosReply, error) {
	log.Printf("GetAllTodos: Attempting to fetch all todos")

	todos, err := s.Queries.GetAllTodos(ctx)

	if err != nil {
		return nil, errors.New("failed to fetch todos")
	}

	var todoReplies []*pb.TodoReply

	for _, todo := range todos {
		todoReply := pb.TodoReply{Id: todo.ID, Title: todo.Title, Completed: todo.Completed, CreatedAt: timestamppb.New(todo.Createdat)}

		todoReplies = append(todoReplies, &todoReply)
	}

	log.Printf("GetAllTodos: Successfully fetched all todos")

	return &pb.GetAllTodosReply{Todos: todoReplies}, nil
}
