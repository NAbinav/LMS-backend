package helper

import (
	"context"
	"dbms/db"
	"fmt"

	// "fmt"

	"github.com/gin-gonic/gin"
)

func CheckRole(c *gin.Context, role string) bool {
	user, err := WhoamI(c)
	if err != nil {
		c.JSON(400, err)
	}
	return user.Role == role

}

func CheckValidFaculty(ctx context.Context, user_id int, course_id int) bool {
	query := "SELECT instructor_id from courses where id=$1;"
	var inst_id int
	err := db.DB.QueryRow(ctx, query, course_id).Scan(&inst_id)
	fmt.Println(inst_id)
	if err != nil {
		return false
	}
	return user_id == inst_id
}
