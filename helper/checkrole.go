package helper

import (
	"context"
	"dbms/db"

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
	query := "SELECT instructor_id from courses where id=$2"
	rows, err := db.DB.Query(ctx, query, user_id, course_id)
	if err != nil {
		return false
	}
	var inst_id int
	defer rows.Close()
	rows.Scan(&inst_id)
	return user_id == inst_id
}
