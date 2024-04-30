package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	r := gin.Default()

	// create group student
	student := r.Group("/students")

	// set middleware basic auth
	student.Use(gin.BasicAuth(gin.Accounts{
		"admin": "admin",
		"user":  "user",
	}))

	student.Use(func(c *gin.Context) {
		// do something

		// log ip
		ip := c.ClientIP()
		log.Println("IP:", ip)

		c.Next()
	})

	// define routes
	student.GET("/", studentHandler.GetStudents)
	student.GET("/:id", studentHandler.GetStudentByID)
	student.POST("/", studentHandler.CreateStudent)
	student.PUT("/:id", studentHandler.UpdateStudent)
	student.DELETE("/:id", studentHandler.DeleteStudent)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	r.Run(":8080")
}
