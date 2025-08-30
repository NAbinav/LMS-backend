package enrollment

import (
	"context"
	"dbms/db"
	"fmt"
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
