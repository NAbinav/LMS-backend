package submission

import (
	"context"
	"dbms/db"
	"fmt"
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
