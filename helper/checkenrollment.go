package helper

import (
	"context"
	"dbms/db"
)

func CheckIfEnrolled(ctx context.Context, userID int, courseID string) bool {
	query := `SELECT instructor_id FROM courses WHERE id=$1`
	rows, err := db.DB.Query(ctx, query, courseID)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var instructorID int
		if err := rows.Scan(&instructorID); err != nil {
			return false
		}
		if instructorID == userID {
			return true
		}
	}

	return false
}

