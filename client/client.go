package main

import (
	"context"
	"fmt"
	"log"

	studentpb "student-service/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := studentpb.NewStudentServiceClient(conn)

	res, err := client.GetStudent(context.Background(), &studentpb.GetStudentRequest{Id: 1})
	if err != nil {
		log.Fatalf("error getting student: %v", err)
	}

	fmt.Println("Student Info:")
	fmt.Println("ID:", res.Student.Id)
	fmt.Println("Name:", res.Student.Name)
}
