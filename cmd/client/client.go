package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	studentpb "student-service/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := studentpb.NewStudentServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Student Menu ---")
		fmt.Println("1. List Students")
		fmt.Println("2. Add Student")
		fmt.Println("3. Update Student")
		fmt.Println("4. Delete Student")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			listStudents(client)
		case "2":
			addStudent(client, reader)
		case "3":
			updateStudent(client, reader)
		case "4":
			deleteStudent(client, reader)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func listStudents(client studentpb.StudentServiceClient) {
	res, err := client.ListStudents(context.Background(), &studentpb.ListStudentsRequest{})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("\nStudents:")
	for _, s := range res.Students {
		fmt.Printf("ID: %d, Name: %s\n", s.Id, s.Name)
	}
}

func addStudent(client studentpb.StudentServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	res, err := client.CreateStudent(context.Background(), &studentpb.CreateStudentRequest{Name: name})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Created Student:", res.Student.Id, res.Student.Name)
}

func updateStudent(client studentpb.StudentServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter student ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	res, err := client.UpdateStudent(context.Background(), &studentpb.UpdateStudentRequest{
		Id:   int32(id),
		Name: name,
	})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Updated Student:", res.Student.Id, res.Student.Name)
}

func deleteStudent(client studentpb.StudentServiceClient, reader *bufio.Reader) {
	fmt.Print("Enter student ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, _ := strconv.Atoi(idStr)

	_, err := client.DeleteStudent(context.Background(), &studentpb.DeleteStudentRequest{
		Id: int32(id),
	})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Deleted student with ID:", id)
}
