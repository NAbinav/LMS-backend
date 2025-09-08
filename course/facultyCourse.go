package course

import (
	"dbms/helper"
	"fmt"
	// "errors"

	"github.com/gin-gonic/gin"
)

func GetFacultyCourses(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(401, err)
	}
	all_courses, err := AllCoursesHandled(ctx, user.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
		return
	}
	c.JSON(200, all_courses)
}
