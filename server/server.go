package server

import (
	"github.com/jamiebmurray25/grpc-crud/database"
	pb "github.com/jamiebmurray25/grpc-crud/protobuf"
)

type Server struct {
	pb.UnimplementedTodoServer
	Queries *database.Queries
}
