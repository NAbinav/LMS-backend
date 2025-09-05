package enrollment

import (
	"dbms/helper"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EnrollUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := helper.WhoamI(c)
	if err != nil {
		fmt.Println("hello")
		c.JSON(401, "Token processing problem")
		return
	}
	var course_id int
	course_id_str := c.Query("course")
	if course_id_str == "" {
		c.JSON(400, fmt.Errorf("course_id is missing"))
		return
	}
	fmt.Println(course_id)
	course_id, err = strconv.Atoi(course_id_str)
	if err != nil {
		c.JSON(400, err)
		return
	}
	err = EnrollUser(ctx, user.Id, course_id)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, "Enrolled properly")
	c.Abort()
}

func GetEnrolled(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(401, "Token Unauthorised")
	}
	if user.Role != "student" {
		c.JSON(200, errors.New("No enrolling for faculties"))
		return
	}
	all_courses, err := GetAllEnrolledCourse(ctx, user.Id)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, all_courses)
}

func DeleteEnrollementHandler(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := helper.WhoamI(c)
	if err != nil {
		fmt.Println("hello")
		c.JSON(401, "Token processing problem")
		return
	}
	var course_id int
	course_id_str := c.Query("course")
	if course_id_str == "" {
		c.JSON(400, fmt.Errorf("course_id is missing"))
		return
	}
	fmt.Println(course_id)
	course_id, err = strconv.Atoi(course_id_str)
	if err != nil {
		c.JSON(400, err)
		return
	}
	err = DeleteEnrollment(ctx, course_id, user.Id)
	if err != nil {
		c.JSON(400, err)
		return
	}
	fmt.Println(err)
	// c.JSON(200, )

}
