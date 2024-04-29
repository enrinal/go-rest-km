package repository

import (
	"errors"
	"rest/model"
)

var studentDummy = []model.Student{
	{
		ID:    1,
		Name:  "John Doe",
		Age:   20,
		Email: "john.doe@student.itb.ac.id",
	},
	{
		ID:    2,
		Name:  "Jane Doe",
		Age:   21,
		Email: "jane.doe@student.itb.ac.id",
	},
	{
		ID:    3,
		Name:  "John Smith",
		Age:   22,
		Email: "john.smith@student.itb.ac.id",
	},
}

type StudentRepository struct {
}

type StudentRepositoryInterface interface {
	GetStudentByID(id int) (*model.Student, error)
	GetStudents() ([]model.Student, error)
	UpdateStudent(student *model.Student) (*model.Student, error)
	DeleteStudent(id int) error
	CreateStudent(student *model.Student) (*model.Student, error)
}

// NewStudentRepository creates a new instance of StudentRepository
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

// GetStudentByID returns a student by its ID
func (r *StudentRepository) GetStudentByID(id int) (*model.Student, error) {
	for _, s := range studentDummy {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, errors.New("student not found")
}

// GetStudents returns all students
func (r *StudentRepository) GetStudents() ([]model.Student, error) {
	return studentDummy, nil
}

// UpdateStudent updates a student
func (r *StudentRepository) UpdateStudent(student *model.Student) (*model.Student, error) {
	for i, s := range studentDummy {
		if s.ID == student.ID {
			studentDummy[i] = *student
			return student, nil
		}
	}
	return nil, errors.New("student not found")
}

// DeleteStudent deletes a student by its ID
func (r *StudentRepository) DeleteStudent(id int) error {
	for i, s := range studentDummy {
		if s.ID == id {
			studentDummy = append(studentDummy[:i], studentDummy[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}

// CreateStudent creates a new student
func (r *StudentRepository) CreateStudent(student *model.Student) (*model.Student, error) {
	studentDummy = append(studentDummy, *student)
	return student, nil
}
