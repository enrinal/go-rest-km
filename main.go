package main

import (
	"log"
	"net/http"
	"rest/handler"
	"rest/repository"
	"rest/service"
)

func main() {
	// repository
	studentRepo := repository.NewStudentRepository()
	// service
	studentService := service.NewStudentService(studentRepo)
	// handler
	studentHandler := handler.NewStudentHandler(studentService)

	// start the server
	mux := http.NewServeMux()
	mux.HandleFunc("/student", studentHandler.StudentHandler)
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", mux)
}
