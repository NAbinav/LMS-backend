package submission

import (
	"context"
	"dbms/db"
	"fmt"
	"time"
)

func CheckIfAssigned(ctx context.Context, user_id int, assignment_id int) bool {
	query := "select e.user_id from assignments a join enrollments e on a.course_id=e.course_id where a.id=$1 and user_id=$2"
	fmt.Println(query)
	var query_user int
	err := db.DB.QueryRow(ctx, query, assignment_id, user_id).Scan(&query_user)
	fmt.Println(query_user)
	if err != nil {
		return false
	}
	return user_id == query_user
}

func AssignmentSubmit(ctx context.Context, user_id int, a_id int, submission_text string) error {
	query := "insert into submissions (user_id,assignment_id,submission_text,submitted_at) values ($1,$2,$3,$4)"
	_, err := db.DB.Exec(ctx, query, user_id, a_id, submission_text, time.Now())
	return err
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
