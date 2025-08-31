package helper

import (
	"dbms/db"
	"dbms/jwt"
	"dbms/schema"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WhoamI(c *gin.Context) (schema.User, error) {
	ctx := c.Request.Context()
	token, err := c.Cookie("token")
	fmt.Println(token)
	if err != nil {
		return schema.User{}, fmt.Errorf("token invalid")
	}
	email, err := jwt.Verify_JWT(token)
	fmt.Println(email, err)
	if err != nil {
		return schema.User{}, fmt.Errorf("Invalid Token")
	}
	query := "SELECT * FROM users WHERE email = $1"

	var user schema.User

	row := db.DB.QueryRow(ctx, query, email)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
		fmt.Println("Error fetching user:", err)
		return schema.User{}, fmt.Errorf("Error fetching user")
	}

	return user, nil

}
