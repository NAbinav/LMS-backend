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
	"dbms/helper"

	"github.com/gin-gonic/gin"
)

func CreateQuiz(ctx context.Context, course_id int, title string, max_attempts int, time_limit int) error {
	query := "insert into quizzes (course_id,title,max_attempts,time_limit) values ($1,$2,$3,$4)"
	_, err := db.DB.Exec(ctx, query, course_id, title, max_attempts, time_limit)
	if err != nil {
		return err
	}
	return nil
}

func CheckQuizEnrolled(c *gin.Context, quiz_id int) bool {
	user, err := helper.WhoamI(c)
	if err != nil {
		return false
	}
	query := "SELECT * FROM QUIZZES JOIN COURSES c ON c.id=q.course"

	return true
}
