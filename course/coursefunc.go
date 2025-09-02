package course

import (
	"context"
	"dbms/db"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Course_type struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	InstructorID int    `json:"instructor_id"`
	Credits      int    `json:"credits"`
}
type Courses struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	InstructorID int       `json:"instructor_id"`
	Credits      int       `json:"credits"`
	CreatedAt    time.Time `json:"created_at"`
}

func CreateCourseDB(ctx context.Context, id int, title string, description string, instructorID int, credits int) error {
	query := "INSERT INTO courses (title, description, instructor_id, credits) VALUES ($1, $2, $3, $4)"
	_, err := db.DB.Exec(ctx, query, title, description, instructorID, credits)
	if err != nil {
		fmt.Println("Error creating course:", err)
		return err
	}
	return nil

}

func CreateCourse(c *gin.Context) {
	ctx := c.Request.Context()
	var Course Course_type
	err := c.ShouldBindJSON(&Course)
	if err != nil {
		c.JSON(400, "Provide body")
	}
	err = CreateCourseDB(ctx, Course.Id, Course.Title, Course.Description, Course.InstructorID, Course.Credits)
	if err != nil {
		fmt.Println("Error creating course:", err)
		c.JSON(500, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(200, gin.H{"message": "Course created successfully"})
}

func DeleteCourse(c *gin.Context) {
	id := c.Query("id")
	ctx := c.Request.Context()
	query := "DELETE FROM courses WHERE id = $1"
	result, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		fmt.Println("Error deleting course:", err)
		c.JSON(500, gin.H{"error": "Failed to delete course"})
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Course deleted successfully"})
}

func GetCourse(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	query := "SELECT * FROM courses WHERE id = $1"
	row := db.DB.QueryRow(ctx, query, id)

	var course struct {
		Id           int       `json:"id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		InstructorID int       `json:"instructor_id"`
		Credits      int       `json:"credits"`
		CreatedAt    time.Time `json:"created_at"`
	}

	if err := row.Scan(&course.Id, &course.Title, &course.Description, &course.InstructorID, &course.Credits, &course.CreatedAt); err != nil {
		fmt.Println("Error fetching course:", err)
		c.JSON(404, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(200, course)
}

func ListCourses(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Query("inst_id")
	query := "SELECT * FROM courses"
	rows, err := db.DB.Query(ctx, query, id)
	if err != nil {
		fmt.Println("Error fetching courses:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch courses"})
		return
	}
	defer rows.Close()

	var courses []Courses

	for rows.Next() {
		var course Courses
		if err := rows.Scan(&course.Id, &course.Title, &course.Description, &course.InstructorID, &course.Credits, &course.CreatedAt); err != nil {
			fmt.Println("Error scanning course row:", err)
			c.JSON(500, gin.H{"error": "Failed to scan course row"})
			return
		}
		courses = append(courses, course)
	}

	c.JSON(200, courses)
}

func UpdateCourse(c *gin.Context) {
	var course Course_type
	if err := c.ShouldBindJSON(&course); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	ctx := c.Request.Context()
	query := "UPDATE courses SET title = $1, description = $2, instructor_id = $3, credits = $4 WHERE id = $5"
	fmt.Println("Executing query:", query, course.Title, course.Description, course.InstructorID, course.Credits, course.Id)
	result, err := db.DB.Exec(ctx, query, course.Title, course.Description, course.InstructorID, course.Credits, course.Id)
	if err != nil {
		fmt.Println("Error updating course:", err)
		c.JSON(500, gin.H{"error": "Failed to update course"})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Course updated successfully"})
}
