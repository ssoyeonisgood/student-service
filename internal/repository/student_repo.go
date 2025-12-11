package repository

import (
	"student-service/internal/db"
	pb "student-service/pb"
)

func GetStudent(id int32) (*pb.Student, error) {
	row := db.DB.QueryRow("SELECT id, name FROM students WHERE id = $1", id)

	var student pb.Student
	err := row.Scan(&student.Id, &student.Name)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func CreateStudent(name string) (*pb.Student, error) {
	var id int32

	err := db.DB.QueryRow(
		"INSERT INTO students (name) VALUES ($1) RETURNING id",
		name,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.Student{
		Id:   id,
		Name: name,
	}, nil
}

func UpdateStudent(id int32, name string) (*pb.Student, error) {
	_, err := db.DB.Exec(
		"UPDATE students SET name=$1 WHERE id=$2",
		name,
		id,
	)
	if err != nil {
		return nil, err
	}

	return &pb.Student{
		Id:   id,
		Name: name,
	}, nil
}

func DeleteStudent(id int32) error {
	_, err := db.DB.Exec("DELETE FROM students WHERE id=$1", id)
	return err
}

func ListStudents() ([]*pb.Student, error) {
	rows, err := db.DB.Query("SELECT id, name FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*pb.Student

	for rows.Next() {
		var s pb.Student
		rows.Scan(&s.Id, &s.Name)
		students = append(students, &s)
	}

	return students, nil
}
