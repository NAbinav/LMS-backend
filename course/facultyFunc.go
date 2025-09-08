package course

import (
	"context"
	"dbms/db"
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
