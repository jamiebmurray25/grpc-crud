FROM golang:1.22.5 AS build

WORKDIR /grpc-crud

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o grpc-crud .
RUN go build -o ./sql/seed ./sql/seed.go

FROM debian:latest

RUN apt update && apt upgrade
RUN apt install sqlite3 -y

WORKDIR /grpc-crud/

COPY --from=build /grpc-crud/grpc-crud .
COPY --from=build /grpc-crud/sql/seed /grpc-crud/sql/schema.sql ./sql/

RUN ./sql/seed

EXPOSE 3333

ENTRYPOINT [ "./grpc-crud" ]