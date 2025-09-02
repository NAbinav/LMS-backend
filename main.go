package main

import (
	"dbms/course"
	"dbms/db"
	"dbms/enrollment"
	"dbms/handler"
	"dbms/quiz"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	if err := db.Initiate_DB(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	r.GET("/role", handler.ListUserFromRole)
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.LoginHandler)
	r.GET("/get_user", handler.Getuser)
	r.DELETE("/user", handler.DeleteUser)

	r.GET("/course", course.GetCourse)
	r.POST("/course", course.CreateCourse)
	r.PUT("/course", course.UpdateCourse)
	r.DELETE("/course", course.DeleteCourse)

	r.GET("/allcourse", course.ListCourses)

	r.GET("/enroll", enrollment.GetEnrolled)
	r.POST("/enroll", enrollment.EnrollUserHandler)

	r.GET("/quiz", quiz.CreateQuizHandler)
	r.POST("/quiz", quiz.CreateQuizHandler)
	r.Run(":8080")
}
