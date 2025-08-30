package user

import (
	"context"
	"dbms/db"
	"fmt"
)


func RegisterUser(ctx context.Context, name string, email string, password string, role string) error {
	query := "INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)"
	fmt.Println(query, name, email, password, role) 
	_, err := db.DB.Exec(ctx, query, name, email, password, role)
	return err
}

func LoginUser(ctx context.Context, email string, password string) (int,string,error) {
	query := "SELECT id,role FROM users WHERE email = $1 AND password = $2"
	var id int
	var role string // dv change
	
	err := db.DB.QueryRow(ctx, query, email, password).Scan(&id,&role)
	if err != nil {
		return 0,"", err
	}
	return id,role,nil
}
