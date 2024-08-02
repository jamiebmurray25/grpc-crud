package server

import (
	"context"
	"errors"
	"log"

	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) GetTodoById(ctx context.Context, in *pb.GetTodoByIdRequest) (*pb.TodoReply, error) {
	log.Printf("Received: %v", in.GetId())

	todo, err := s.Queries.GetTodoById(ctx, in.GetId())

	log.Printf("%+v\n", todo)

	if err != nil {
		return nil, errors.New("todo not found")
	}

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title, Completed: todo.Completed, CreatedAt: timestamppb.New(todo.Createdat)}, nil
}
