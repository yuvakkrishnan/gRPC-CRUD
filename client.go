package main

import (
	"context"
	"log"

	pb "github.com/yuvakkrishnan/gRPC-CRUD/grpc-crud/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Create user
	user := &pb.User{Id: 1, Name: "John Doe", Email: "john@example.com"}
	createResp, err := c.CreateUser(context.Background(), &pb.CreateUserRequest{User: user})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Created user: %v", createResp.GetUser())

	// Get user
	getResp, err := c.GetUser(context.Background(), &pb.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("Got user: %v", getResp.GetUser())

	// Update user
	user.Name = "John Updated"
	updateResp, err := c.UpdateUser(context.Background(), &pb.UpdateUserRequest{User: user})
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}
	log.Printf("Updated user: %v", updateResp.GetUser())

	// Delete user
	deleteResp, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	log.Printf("Deleted user successfully: %v", deleteResp.GetSuccess())
}
