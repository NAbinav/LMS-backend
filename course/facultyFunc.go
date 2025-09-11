package course

import (
	"context"
	"dbms/db"
	"fmt"
)

func AllCoursesHandled(ctx context.Context, instructor_id int) ([]Courses, error) {
	query := "SELECT id,title,description,credits,created_at FROM courses where instructor_id=$1"
	rows, err := db.DB.Query(ctx, query, instructor_id)
	if err != nil {
		return []Courses{}, err
	}
	var AllCourse []Courses
	defer rows.Close()
	for rows.Next() {
		var SingleCourse Courses
		if err := rows.Scan(&SingleCourse.Id, &SingleCourse.Title, &SingleCourse.Description, &SingleCourse.Credits, &SingleCourse.CreatedAt); err != nil {
			return []Courses{}, err
		}
		AllCourse = append(AllCourse, SingleCourse)
	}
	return AllCourse, nil
}

func AllStudentsEnrolled(ctx context.Context, course_id string) ([]string, error) {
	query := "SELECT u.name FROM ENROLLMENTS e join users u on u.id=e.user_id where course_id=$1"
	rows, err := db.DB.Query(ctx, query, course_id)
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	defer rows.Close()
	var AllNames []string
	for rows.Next() {
		var SingleName string
		if err := rows.Scan(&SingleName); err != nil {
			return []string{}, err
		}
		AllNames = append(AllNames, SingleName)
	}
	return AllNames, nil
}
