package service

import (
	"errors"
	"rest/model"
	"rest/repository"
)

type StudentService struct {
	// dependency injection
	studentRepo *repository.StudentRepository
}

type StudentServiceInterface interface {
	GetStudentByID(id int) (*model.Student, error)
	GetStudents() ([]model.Student, error)
	UpdateStudent(student *model.Student) (*model.Student, error)
	DeleteStudent(id int) error
	CreateStudent(student *model.Student) (*model.Student, error)
}

// NewStudentService creates a new instance of StudentService
func NewStudentService(studentRepo *repository.StudentRepository) *StudentService {
	return &StudentService{studentRepo}
}

// GetStudentByID returns a student by its ID
func (s *StudentService) GetStudentByID(id int) (*model.Student, error) {
	return s.studentRepo.GetStudentByID(id)
}

// GetStudents returns all students
func (s *StudentService) GetStudents() ([]model.Student, error) {
	return s.studentRepo.GetStudents()
}

// UpdateStudent updates a student
func (s *StudentService) UpdateStudent(student *model.Student) (*model.Student, error) {
	// get all students
	students, err := s.studentRepo.GetStudents()
	if err != nil {
		return nil, err
	}

	// check email uniqueness
	for _, s := range students {
		if s.Email == student.Email && s.ID != student.ID {
			return nil, errors.New("email already exists")
		}
	}

	return s.studentRepo.UpdateStudent(student)
}

// DeleteStudent deletes a student by its ID
func (s *StudentService) DeleteStudent(id int) error {
	return s.studentRepo.DeleteStudent(id)
}

// CreateStudent creates a new student
func (s *StudentService) CreateStudent(student *model.Student) (*model.Student, error) {

	// get all students
	students, err := s.studentRepo.GetStudents()
	if err != nil {
		return nil, err
	}

	// check email uniqueness
	for _, s := range students {
		if s.Email == student.Email {
			return nil, errors.New("email already exists")
		}
	}

	// get latest student ID
	lastStudent := students[len(students)-1]
	student.ID = lastStudent.ID + 1

	return s.studentRepo.CreateStudent(student)
}
