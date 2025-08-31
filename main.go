package main

import (
	"dbms/course"
	"dbms/db"
	"dbms/handler"
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
	r.POST("/register", handler.RegisterUser)
	r.POST("/login", handler.LoginHandler)
	r.GET("/get_user", handler.Getuser)
	r.GET("/role", handler.ListUserFromRole)
	r.DELETE("/deleteuser", handler.DeleteUser)
	r.POST("/createcourse", course.CreateCourse)
	r.DELETE("/deletecourse", course.DeleteCourse)
	r.GET("/getcourse", course.GetCourse)
	r.GET("/listcourses", course.ListCourses)
	r.PUT("/updatecourse", course.UpdateCourse)
	r.POST("/enroll", handler.EnrollUserHandler)
	r.GET("/getenrolled", handler.GetEnrolled)
	r.Run(":8080")
}
