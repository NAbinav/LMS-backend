package submission

import (
	"context"
	"dbms/db"
	"fmt"
)

func CheckIfAssigned(ctx context.Context, user_id int, assignment_id int) {
	query := "select * from assignment a join enrollment on a.course_id=e.course_id where e.user_id=$1"
	fmt.Println(query)
	rows, err := db.DB.Query(ctx, query, user_id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {

	}
}
