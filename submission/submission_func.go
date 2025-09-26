package submission

import (
	"context"
	"dbms/db"
	"errors"
	"fmt"
	"time"
)

func CheckIfAssigned(ctx context.Context, user_id int, assignment_id int) bool {
	query := "select e.user_id from assignments a join enrollments e on a.course_id=e.course_id where a.id=$1 and e.user_id=$2"
	fmt.Println(query)
	var query_user int
	err := db.DB.QueryRow(ctx, query, assignment_id, user_id).Scan(&query_user)
	fmt.Println(query_user)
	if err != nil {
		return false
	}
	return user_id == query_user
}

func AssignmentSubmit(ctx context.Context, userID int, assignmentID int, submissionText string) error {
	// 1. Get assignment due_date
	var dueDate time.Time
	err := db.DB.QueryRow(ctx, `SELECT due_date FROM assignments WHERE id = $1`, assignmentID).Scan(&dueDate)
	if err != nil {
		return fmt.Errorf("failed to fetch due date: %w", err)
	}

	now := time.Now()

	query := `INSERT INTO submissions (user_id, assignment_id, submission_text, submitted_at, status) VALUES ($1, $2, $3, $4, $5)`
	status := "submitted"
	if now.After(dueDate) {
		status = "late"
	}
	_, err = db.DB.Exec(ctx, query, userID, assignmentID, submissionText, now, status)
	return err
}

func GradeAssignment_db(ctx context.Context, submissionID int, grade int, teacherID int) error {
	query := `UPDATE submissions SET grade = $1, graded_at = $2, graded_by = $3 WHERE id = $4;`
	result, err := db.DB.Exec(ctx, query, grade, time.Now(), teacherID, submissionID)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// Check if any row was actually updated
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("submission not found")
	}

	return nil
}

func CourseidOfAssignment(ctx context.Context, a_id string) (int, error) {
	query := "select course_id from assignments where id=$1"
	var c_id int
	err := db.DB.QueryRow(ctx, query, a_id).Scan(&c_id)
	if err != nil {
		return 0, err
	}
	return c_id, nil
}

type CustomSubmission struct {
	UserID      int       `json:"user_id"`
	URL         string    `json:"url"`
	SubmittedAt time.Time `json:"submitted_at"`
}

func AllSubmissions(ctx context.Context, a_id string) ([]CustomSubmission, error) {
	query := "select user_id,submission_text,submitted_at  from submissions where assignment_id=$1"
	rows, err := db.DB.Query(ctx, query, a_id)

	if err != nil {
		return []CustomSubmission{}, err
	}
	var Submissions []CustomSubmission
	defer rows.Close()
	for rows.Next() {
		var SingleSubmission CustomSubmission
		if err := rows.Scan(&SingleSubmission.UserID, &SingleSubmission.URL, &SingleSubmission.SubmittedAt); err != nil {
			return []CustomSubmission{}, err
		}
		Submissions = append(Submissions, SingleSubmission)
	}
	return Submissions, err
}

func HasSubmitted(ctx context.Context, userID int, assignmentID int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM submissions WHERE user_id = $1 AND assignment_id = $2)"
	err := db.DB.QueryRow(ctx, query, userID, assignmentID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

type Submission struct {
	AssignmentID int        `json:"assignment_id"`
	SubmittedAt  time.Time  `json:"submitted_at"`
	Grade        *float64   `json:"grade,omitempty"`
	GradedAt     *time.Time `json:"graded_at,omitempty"`
	GradedBy     *int       `json:"graded_by,omitempty"`
}

func GetUserSubmissions(ctx context.Context, userID int) ([]Submission, error) {
	rows, err := db.DB.Query(ctx,
		`SELECT assignment_id, submitted_at, grade, graded_at, graded_by
		 FROM submissions
		 WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submissions []Submission
	for rows.Next() {
		var s Submission
		if err := rows.Scan(&s.AssignmentID, &s.SubmittedAt, &s.Grade, &s.GradedAt, &s.GradedBy); err != nil {
			return nil, err
		}
		submissions = append(submissions, s)
	}
	return submissions, nil
}
