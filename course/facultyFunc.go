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

type Details struct {
	Name  string
	Email string
}

func AllStudentsEnrolled(ctx context.Context, course_id string) ([]Details, error) {
	query := "SELECT u.name,u.email FROM ENROLLMENTS e join users u on u.id=e.user_id where course_id=$1"
	rows, err := db.DB.Query(ctx, query, course_id)
	if err != nil {
		fmt.Println(err)
		return []Details{}, err
	}
	fmt.Println(rows)
	defer rows.Close()
	var AllNames []Details
	for rows.Next() {
		var Name string
		var Email string
		if err := rows.Scan(&Name, &Email); err != nil {
			return []Details{}, err
		}
		var detail struct {
			Name  string
			Email string
		}
		detail.Name = Name
		detail.Email = Email
		AllNames = append(AllNames, detail)
	}
	return AllNames, nil
}
