package server

import (
	"context"
	"errors"
	"log"

	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.TodoReply, error) {

	if in.Completed {
		log.Printf("UpdateTodo: Attempting to update todo '%s' with values: {title: %s, completed: true}", in.GetId(), in.GetTitle())
	} else {
		log.Printf("UpdateTodo: Attempting to update todo '%s' with values: {title: %s, completed: false}", in.GetId(), in.GetTitle())
	}

	todo, err := s.Queries.UpdateTodo(ctx, database.UpdateTodoParams{Title: in.GetTitle(), Completed: in.GetCompleted(), ID: in.GetId()})

	if err != nil {
		return nil, errors.New("failed to update todo")
	}

	log.Printf("UpdateTodo: Successfully updated todo '%s'", in.GetId())

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title, Completed: todo.Completed, CreatedAt: timestamppb.New(todo.Createdat)}, nil
}
