package assignments

import (
	"context"
	"dbms/db"
	"fmt"
	"time"
)

func CreateAssignment(ctx context.Context, course_id int, title string, description string, due_date time.Time, max_points int) error {

	query := "INSERT INTO assignments (course_id, title, description, due_date, max_points) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.DB.Exec(ctx, query, course_id, title, description, due_date, max_points)
	if err != nil {
		fmt.Println("DB error:", err)
	}
	return err
}

type CustomAssignment struct {
	Assi_id     int       `json:"id"`
	Course_name string    `json:"course_name"`
	Assgn_title string    `json:"assignment_title"`
	Due_date    time.Time `json:"due_date"`
	Max_points  int       `json:"max_points"`
	User_name   string    `json:"teacher_name"`
	Description string    `json:"description"`
}

func GetAssignmentFac(ctx context.Context, user_id int) ([]CustomAssignment, error) {
	query := `
	SELECT a.id AS ass_id,
	       c.title AS course_name,
	       a.title AS assignment_title,
	       a.description AS description,
	       a.due_date,
	       a.max_points,
	       u.name AS teacher_name
	FROM assignments a
	JOIN courses c ON c.id = a.course_id
	JOIN users u ON u.id = c.instructor_id
	WHERE u.id = $1;
	`

	rows, err := db.DB.Query(ctx, query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allAssignments []CustomAssignment
	for rows.Next() {
		var ass CustomAssignment
		if err := rows.Scan(
			&ass.Assi_id,
			&ass.Course_name,
			&ass.Assgn_title,
			&ass.Description,
			&ass.Due_date,
			&ass.Max_points,
			&ass.User_name,
		); err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}
		allAssignments = append(allAssignments, ass)
	}
	return allAssignments, nil
}

func GetAssignment(ctx context.Context, user_id int) ([]CustomAssignment, error) {
	query := `
	SELECT a.id AS ass_id,
	       c.title AS course_name,
	       a.title AS assignment_title,
	       a.description AS description,
	       a.due_date,
	       a.max_points,
	       u.name AS teacher_name
	FROM assignments a
	JOIN enrollments e ON e.course_id = a.course_id
	JOIN users u ON u.id = e.user_id
	JOIN courses c ON c.id = e.course_id
	WHERE u.id = $1;
	`

	rows, err := db.DB.Query(ctx, query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allAssignments []CustomAssignment
	for rows.Next() {
		var ass CustomAssignment
		if err := rows.Scan(
			&ass.Assi_id,
			&ass.Course_name,
			&ass.Assgn_title,
			&ass.Description,
			&ass.Due_date,
			&ass.Max_points,
			&ass.User_name,
		); err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}
		allAssignments = append(allAssignments, ass)
	}
	return allAssignments, nil
}

type CustomSubmission struct {
	Submission_id    int       `json:"submission_id"`
	Student_name     string    `json:"student_name"`
	Course_name      string    `json:"course_name"`
	Assignment_title string    `json:"assignment_title"`
	Submission_text  string    `json:"submission_text"`
	Submitted_at     time.Time `json:"submitted_at"`
	Grade            float64   `json:"grade"`
	Status           string    `json:"status"`
}

func GetSubmissionsByAssignment(ctx context.Context, assignmentID int) ([]CustomSubmission, error) {
	query := `
	SELECT 
		s.id AS submission_id,
		u.name AS student_name,
		c.title AS course_name,
		a.title AS assignment_title,
		s.submission_text,
		s.submitted_at,
		COALESCE(s.grade, 0) AS grade,
		s.status
	FROM submissions s
	JOIN assignments a ON a.id = s.assignment_id
	JOIN courses c ON c.id = a.course_id
	JOIN users u ON u.id = s.user_id
	WHERE s.assignment_id = $1 AND u.role = 'student';
	`

	rows, err := db.DB.Query(ctx, query, assignmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submissions []CustomSubmission
	for rows.Next() {
		var sub CustomSubmission
		if err := rows.Scan(
			&sub.Submission_id,
			&sub.Student_name,
			&sub.Course_name,
			&sub.Assignment_title,
			&sub.Submission_text,
			&sub.Submitted_at,
			&sub.Grade,
			&sub.Status,
		); err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}
		submissions = append(submissions, sub)
	}

	return submissions, nil
}
