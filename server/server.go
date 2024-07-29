package server

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedTodoServer
	Queries *database.Queries
}

func (s *Server) GetAllTodos(ctx context.Context, in *pb.Empty) (*pb.GetAllTodosReply, error) {
	todos, err := s.Queries.GetAllTodos(ctx)

	if err != nil {
		return nil, errors.New("Failed to fetch todos.")
	}

	var todoReplies []*pb.TodoReply
	for _, todo := range todos {

		todoReply := pb.TodoReply{Id: todo.ID, Title: todo.Title, CreatedAt: timestamppb.New(todo.Createdat.Time)}

		todoReplies = append(todoReplies, &todoReply)
	}

	return &pb.GetAllTodosReply{Todos: todoReplies}, nil
}

func (s *Server) GetTodoById(ctx context.Context, in *pb.GetTodoByIdRequest) (*pb.TodoReply, error) {
	log.Printf("Received: %v", in.GetId())

	todo, err := s.Queries.GetTodoById(ctx, in.GetId())

	log.Printf("%+v\n", todo)

	if err != nil {
		return nil, errors.New("Todo not found.")
	}

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title}, nil
}

func (s *Server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.TodoReply, error) {

	todoId, err := uuid.NewV7()

	if err != nil {
		return nil, errors.New("Failed to create todo uuid.")
	}

	log.Printf("Creating todo with uuid: %s", todoId.String())

	todo, err := s.Queries.CreateTodo(ctx, database.CreateTodoParams{ID: todoId.String(), Title: in.Title})

	if err != nil {
		return nil, errors.New("Failed to create todo.")
	}

	return &pb.TodoReply{Id: todo.ID, Title: todo.Title}, nil
}
