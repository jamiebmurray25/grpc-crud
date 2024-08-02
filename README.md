# How to run
Requires Go CLI and SQLite3

`$ git clone https://github.com/jamiebmurray25/grpc-crud`  

`$ cd grpc-crud`  

`$ go run sql/seed.go`  

`$ go build && ./grpc-crud`

## Regenerate Protobuf and GRPC Files
Run `./gen_protobuf.sh` in root (requires protoc and go extension)

## Regenerate SQLC Files
Run `sqlc generate` in root (requires sqlc)
