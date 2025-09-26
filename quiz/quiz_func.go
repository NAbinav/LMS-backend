package quiz

// CREATE TABLE Modules (
//     id INT PRIMARY KEY AUTO_INCREMENT,
//     course_id INT NOT NULL,
//     title VARCHAR(201) NOT NULL,
//     order_num INT NOT NULL,
//     content LONGTEXT,
//     FOREIGN KEY (course_id) REFERENCES Courses(id) ON DELETE CASCADE ON UPDATE CASCADE,
//     INDEX idx_course_order (course_id, order_num),
//     UNIQUE KEY uk_course_order (course_id, order_num)
// );

import (
	"context"
	"dbms/db"
	"fmt"
	"time"
	// "dbms/helper"
	// "dbms/schema"
	//
	// "github.com/gin-gonic/gin"
)

func CreateQuiz(ctx context.Context, course_id int, title string, max_attempts int, time_limit int) error {
	query := "insert into quizzes (course_id,title,max_attempts,time_limit) values ($1,$2,$3,$4)"
	_, err := db.DB.Exec(ctx, query, course_id, title, max_attempts, time_limit)
	if err != nil {
		return err
	}
	return nil
}

type CustomQuizEnrolled struct {
	Course_name string `json:"course_name"`
	Quiz_title  string `json:"quiz_title"`
	Time_limit  int    `json:"time_limit"`
}

func AllQuizEnrolled(ctx context.Context, user_id int) []CustomQuizEnrolled {
	query := "select c.title as course_name, q.title as quiz_title, q.time_limit as time_limit from quizzes q join enrollments e on e.course_id=q.course_id join users u on u.id=e.user_id join courses c on c.id=e.course_id where u.id=$1;"
	rows, err := db.DB.Query(ctx, query, user_id)
	if err != nil {
		return []CustomQuizEnrolled{}
	}
	defer rows.Close()
	var AllQuiz []CustomQuizEnrolled
	for rows.Next() {
		var SingleQuiz CustomQuizEnrolled
		if err := rows.Scan(&SingleQuiz.Course_name, &SingleQuiz.Quiz_title, &SingleQuiz.Time_limit); err != nil {
			return []CustomQuizEnrolled{}
		}
		AllQuiz = append(AllQuiz, SingleQuiz)
	}
	return AllQuiz
}

func GetQuizId(q_id int, user_id int) (CustomQuizEnrolled, error) {

	return CustomQuizEnrolled{}, nil
}

type CustomQuiz struct {
	Quiz_id     int       `json:"id"`
	Course_name string    `json:"course_name"`
	Quiz_title  string    `json:"quiz_title"`
	Time_limit  int       `json:"time_limit"`
	Created_at  time.Time `json:"created_at"`
}

func GetQuizzesByTeacher(ctx context.Context, teacherID int) ([]CustomQuiz, error) {
	query := `
	SELECT 
		q.id AS quiz_id,
		c.title AS course_name,
		q.title AS quiz_title,
		COALESCE(q.time_limit, 0) AS time_limit,
		q.created_at
	FROM quizzes q
	JOIN courses c ON c.id = q.course_id
	WHERE c.instructor_id = $1;
	`

	rows, err := db.DB.Query(ctx, query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []CustomQuiz
	for rows.Next() {
		var q CustomQuiz
		if err := rows.Scan(
			&q.Quiz_id,
			&q.Course_name,
			&q.Quiz_title,
			&q.Time_limit,
			&q.Created_at,
		); err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}
		quizzes = append(quizzes, q)
	}

	return quizzes, nil
}
