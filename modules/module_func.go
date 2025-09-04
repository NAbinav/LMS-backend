package modules

import (
	"context"
	"dbms/db"
)

func CreateModule(ctx context.Context, course_id int, order_num int, title string, content string, link string) error {
	query := "INSERT INTO MODULES (title,content,order_num,course_id,link) values ($1,$2,$3,$4,$5)"
	_, err := db.DB.Exec(ctx, query, title, content, order_num, course_id, link)
	return err
}
