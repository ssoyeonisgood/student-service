package saerver

import (
	"log"
	"net"

	"student-service/internal/db"
	"student-service/internal/service"
	pb "student-service/pb"

	"google.golang.org/grpc"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStudentServiceServer(grpcServer, &service.StudentServiceServer{})

	log.Println("Server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
