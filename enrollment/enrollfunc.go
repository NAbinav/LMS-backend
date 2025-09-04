package enrollment

import (
	"context"
	"dbms/db"
	"dbms/schema"
	"fmt"
	// "github.com/go-playground/locales/qu"
	// "github.com/gin-gonic/gin"
	// "github.com/go-playground/locales/qu"
)

// CREATE TABLE Enrollments (
//     user_id INT NOT NULL,
//     course_id INT NOT NULL,
//     enrollment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     status ENUM('active', 'completed', 'dropped', 'pending') NOT NULL DEFAULT 'active',
//     PRIMARY KEY (user_id, course_id),
//     FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
//     FOREIGN KEY (course_id) REFERENCES Courses(id) ON DELETE CASCADE,
//     INDEX idx_enrollment_date (enrollment_date),
//     INDEX idx_status (status)
// );

func EnrollUser(ctx context.Context, userID int, courseID int) error {
	query := "INSERT INTO Enrollments (user_id, course_id) VALUES ($1, $2)"
	_, err := db.DB.Exec(ctx, query, userID, courseID)
	if err != nil {
		fmt.Println("Error enrolling user:", err)
		return err
	}
	return nil
}

func GetAllEnrolledCourse(ctx context.Context, userID int) ([]schema.Courses, error) {
	query := "SELECT c.id,c.title,c.description,c.instructor_id,c.credits,c.created_at FROM enrollments e join courses c on e.course_id=c.id WHERE e.user_id=$1"
	rows, err := db.DB.Query(ctx, query, userID)
	if err != nil {
		return []schema.Courses{}, err
	}
	defer rows.Close()
	var courses []schema.Courses
	for rows.Next() {
		var course_row schema.Courses
		if err := rows.Scan(&course_row.Id, &course_row.Title, &course_row.Description, &course_row.InstructorID, &course_row.Credits, &course_row.CreatedAt); err != nil {
			return []schema.Courses{}, err
		}
		courses = append(courses, course_row)
	}
	return courses, nil
}

func DeleteEnrollment(ctx context.Context, courseid int, userid int) error {
	query := "DELETE FROM enrollments where user_id=$1 and course_id=$2"
	rows, err := db.DB.Exec(ctx, query, userid, courseid)
	fmt.Println(query, userid, courseid, rows.RowsAffected)
	return err
}
