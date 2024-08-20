package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/yuvakkrishnan/gRPC-CRUD/grpc-crud/user"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users map[int32]*pb.User
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := req.GetUser()
	s.users[user.Id] = user
	return &pb.CreateUserResponse{User: user}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, exists := s.users[req.GetId()]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return &pb.GetUserResponse{User: user}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := req.GetUser()
	s.users[user.Id] = user
	return &pb.UpdateUserResponse{User: user}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, exists := s.users[req.GetId()]
	if !exists {
		return &pb.DeleteUserResponse{Success: false}, nil
	}
	delete(s.users, req.GetId())
	return &pb.DeleteUserResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{users: make(map[int32]*pb.User)})
	log.Printf("server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
