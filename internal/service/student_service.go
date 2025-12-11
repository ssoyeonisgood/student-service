package service

import (
	"context"
	"student-service/internal/repository"
	pb "student-service/pb"
)

type StudentServiceServer struct {
	pb.UnimplementedStudentServiceServer
}

func (s *StudentServiceServer) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	student, err := repository.CreateStudent(req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.CreateStudentResponse{Student: student}, nil
}

func (s *StudentServiceServer) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	student, err := repository.GetStudent(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStudentResponse{Student: student}, nil
}

func (s *StudentServiceServer) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	student, err := repository.UpdateStudent(req.Id, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStudentResponse{Student: student}, nil
}

func (s *StudentServiceServer) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	err := repository.DeleteStudent(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentResponse{Success: true}, nil
}

func (s *StudentServiceServer) ListStudents(ctx context.Context, req *pb.ListStudentsRequest) (*pb.ListStudentsResponse, error) {
	list, err := repository.ListStudents()
	if err != nil {
		return nil, err
	}
	return &pb.ListStudentsResponse{Students: list}, nil
}
