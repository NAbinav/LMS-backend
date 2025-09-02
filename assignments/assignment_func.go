package assignments

import (
	"context"
	"dbms/db"
	"fmt"
	// "github.com/go-playground/locales/qu"
	// "time"
)

func CreateAssignment(ctx context.Context, course_id int, title string, description string, due_date string, max_points int) error {
	// now_time := time.Now()
	// formatted_time := time.Time(now_time).Format(due_date)
	// fmt.Println(formatted_time)
	query := "insert into assignments (course_id,title,description,due_date,max_points) values ($1,$2,$3,$4,$5)"
	_, err := db.DB.Exec(ctx, query, course_id, title, description, due_date, max_points)
	fmt.Println(err)
	return err
}

type CustomAssignment struct {
	course_name string
	assgn_title string
	due_date    string
	max_points  int
}

func GetAssignment(ctx context.Context, user_id int) ([]CustomAssignment, error) {
	query := "select c.title as course_name, a.title as title, a.due_date as due_date ,a.max_points as max_points ,u.name  as user_name from assignments a join enrollments e on e.course_id=a.course_id join users u on u.id=e.user_id join courses c on c.id=e.course_id where u.id=$1;"
	rows, err := db.DB.Query(ctx, query, user_id)
	if err != nil {
		return []CustomAssignment{}, err
	}
	defer rows.Close()
	var all_assignments []CustomAssignment
	for rows.Next() {
		var SingleAssigment CustomAssignment
		if err := rows.Scan(&SingleAssigment.course_name, &SingleAssigment.assgn_title, SingleAssigment.due_date, SingleAssigment.max_points); err != nil {
			return []CustomAssignment{}, err
		}
		all_assignments = append(all_assignments, SingleAssigment)
	}
	return all_assignments, nil
}
