package main

import (
	"context"
	"fmt"
	"log"
	"net"

	studentpb "student-service/proto"

	"google.golang.org/grpc"
)

type StudentServer struct {
	studentpb.UnimplementedStudentServiceServer
}

func (s *StudentServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.GetStudentResponse, error) {
	student := &studentpb.Student{
		Id:   req.Id,
		Name: "Alice",
	}
	return &studentpb.GetStudentResponse{Student: student}, nil
}

func (s *StudentServer) CreateStudent(ctx context.Context, req *studentpb.CreateStudentRequest) (*studentpb.CreateStudentResponse, error) {
	newStudent := &studentpb.Student{
		Id:   1,
		Name: req.Name,
	}
	return &studentpb.CreateStudentResponse{Student: newStudent}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(grpcServer, &StudentServer{})

	fmt.Println("Student Service gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
