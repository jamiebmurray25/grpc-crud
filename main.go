package main

import (
	"context"
	"errors"
	pb "github.com/jamiebmurray25/grpc-crud/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedTodoServer
}

func (s *server) GetTodo(ctx context.Context, in *pb.TodoRequest) (*pb.TodoReply, error) {
	log.Printf("Received: %v", in.GetId())

	if in.GetId() == 1 {
		return &pb.TodoReply{Title: "Clean kitchen"}, nil
	}

	return nil, errors.New("Todo not found.")
}

func main() {
	listener, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	registrar := grpc.NewServer()
	pb.RegisterTodoServer(registrar, &server{})
	reflection.Register(registrar)
	if err := registrar.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
