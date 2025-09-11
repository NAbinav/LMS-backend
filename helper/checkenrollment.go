package helper

import (
	"context"
	"dbms/db"
)

func CheckIfEnrolled(ctx context.Context, user_id int, c_id string) bool {
	query := "select user_id from enrollments where user_id=$1 and course_id=$2"
	var check_user int
	err := db.DB.QueryRow(ctx, query, user_id, c_id).Scan(&check_user)
	if err != nil {
		return false
	}
	if check_user != user_id {
		return false
	}
	return true

}
