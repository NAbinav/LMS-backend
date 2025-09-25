package quizqn

import (
	"context"
	"dbms/db"
	"dbms/schema"
	"fmt"
)

func AddQn_db(ctx context.Context, QuizQns []schema.QuizQuestion) error {
	query := "insert into quizqn values ($1,$2,$3,$4,$5)"
	for i, j := range QuizQns {
		fmt.Println(i, j)
		_, err := db.DB.Exec(ctx, query, j.Quiz_id, j.Question_text, j.Correct_answer, j.Points, j.Order_num)
		if err != nil {
			return err
		}
	}
	return nil
}
