# How to run


### Build from source (requires Go CLI and SQLite3)

`$ git clone https://github.com/jamiebmurray25/grpc-crud`

`$ cd grpc-crud`

`$ go run sql/seed.go`

`$ go build && ./grpc-crud`

### Run with docker

`$ docker build -t jamiebmurray25/grpc-crud .`

`$ docker run -p 3333:3333 localhost/jamiebmurray25/grpc-crud:latest`

## Regenerate Protobuf and GRPC Files
Run `./gen_protobuf.sh` in root (requires [protoc](https://grpc.io/docs/protoc-installation/) and [go extension](https://protobuf.dev/reference/go/go-generated/))

## Regenerate SQLC Files
Run `sqlc generate` in root (requires [sqlc](https://sqlc.dev/))

# Example Usage With [grpcurl](https://github.com/fullstorydev/grpcurl)


`$ grpcurl -plaintext localhost:3333 todo.Todo.GetAllTodos`

`$ grpcurl -plaintext -d '{"id": "<uuid>"}' localhost:3333 todo.Todo.GetTodoById`

`$ grpcurl -plaintext -d '{"id": "<uuid>"}' localhost:3333 todo.Todo.DeleteTodo`

`$ grpcurl -plaintext -d '{"title": "Tidy up the garden"}' localhost:3333 todo.Todo.CreateTodo`

`$ grpcurl -plaintext -d '{"title": "Tidy up the garden", "completed": true, "id": "<uuid>"}' localhost:3333 todo.Todo.UpdateTodo`

