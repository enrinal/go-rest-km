package handler

import (
	"encoding/json"
	"net/http"
	"rest/model"
	"rest/service"
	"strconv"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{studentService}
}

func (h *StudentHandler) StudentHandler(w http.ResponseWriter, r *http.Request) {
	// get student by id
	if r.Method == http.MethodGet {
		if r.URL.Query().Get("id") != "" {
			h.GetStudentByID(w, r)
			return
		}

		h.GetStudents(w, r)
	}

	if r.Method == http.MethodPost {
		h.CreateStudent(w, r)
		return
	}

	if r.Method == http.MethodDelete {
		h.DeleteStudent(w, r)
		return
	}

	if r.Method == http.MethodPut {
		h.UpdateStudent(w, r)
		return
	}
}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.studentService.GetStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studentJson, err := json.Marshal(students)
	if err != nil {
		http.Error(w, "Error while encoding students to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentJson)
	w.WriteHeader(http.StatusOK)
}

func (h *StudentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	// get id from query parameter
	id := r.URL.Query().Get("id")

	// convert id to int
	// if id is not a number, return an error
	studentID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	// call service to get student by id
	student, err := h.studentService.GetStudentByID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// return student as JSON response
	studentJson, err := json.Marshal(student)
	if err != nil {
		http.Error(w, "Error while encoding student to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentJson)
	w.WriteHeader(http.StatusOK)
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	// decode the request body to a student struct
	var student model.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Error while decoding request body", http.StatusInternalServerError)
		return
	}

	// call service to create student
	createdStudent, err := h.studentService.CreateStudent(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return created student as JSON response
	studentJson, err := json.Marshal(createdStudent)
	if err != nil {
		http.Error(w, "Error while encoding student to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentJson)
	w.WriteHeader(http.StatusCreated)
}

func (h *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// get id from query parameter
	id := r.URL.Query().Get("id")

	// convert id to int
	// if id is not a number, return an error
	studentID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	// call service to delete student by id
	err = h.studentService.DeleteStudent(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	// get id from query parameter
	id := r.URL.Query().Get("id")

	// convert id to int
	// if id is not a number, return an error
	studentID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID must be a number", http.StatusBadRequest)
		return
	}

	// decode the request body to a student struct
	var student model.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Error while decoding request body", http.StatusInternalServerError)
		return
	}

	// set the student ID to the ID from the query parameter
	student.ID = studentID

	// call service to update student
	updatedStudent, err := h.studentService.UpdateStudent(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return updated student as JSON response
	studentJson, err := json.Marshal(updatedStudent)
	if err != nil {
		http.Error(w, "Error while encoding student to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentJson)
	w.WriteHeader(http.StatusOK)
}
