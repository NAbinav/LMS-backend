package helper

import "github.com/gin-gonic/gin"

func CheckRole(c *gin.Context, role string) bool {
	user, err := WhoamI(c)
	if err != nil {
		c.JSON(400, err)
	}
	return user.Role == role

}
