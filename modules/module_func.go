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

type OutModule struct {
	Id        int    `json:"id"`
	Course_id int    `json:"course_id"`
	Title     string `json:"title"`
	Order_num int    `json:"order_num"`
	Content   string `json:"content"`
	Link      string `json:"link"`
}

func GetModules(ctx context.Context, mod string, c_id string) (OutModule, error) {
	query := "select * from modules where course_id = $1 and order_num=$2"
	var module OutModule
	err := db.DB.QueryRow(ctx, query, mod, c_id).Scan(&module.Id, &module.Course_id, &module.Title, &module.Order_num, &module.Content, &module.Link)
	if err != nil {
		return OutModule{}, err
	}
	return module, nil
}
