package assignments

import (
	"dbms/helper"
	"fmt"
	// "fmt"

	"github.com/gin-gonic/gin"
)

func CreateAssignmentHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var Assignment struct {
		Course_id   int    `json:"course_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Due_date    string `json:"due_date"`
		Max_points  int    `json:"max_points"`
	}
	err := c.ShouldBindJSON(&Assignment)
	if err != nil {
		c.JSON(400, "Cant create Assignment")
		return
	}
	user, err := helper.WhoamI(c)
	if err != nil || user.Role != "instructor" || !helper.CheckValidFaculty(ctx, user.Id, Assignment.Course_id) {
		c.JSON(401, "Unauthorised Access")
		return
	}
	err = CreateAssignment(ctx, Assignment.Course_id, Assignment.Title, Assignment.Description, Assignment.Due_date, Assignment.Max_points)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(201, "Inserted Successfully")
}

func GetAssignmentHandler(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := helper.WhoamI(c)
	fmt.Println(err)
	if err != nil {
		c.JSON(400, err)
		return
	}
	var assignment []CustomAssignment
	if user.Role != "instructor" {
		assignment, err = GetAssignment(ctx, user.Id)
		fmt.Println(assignment)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, assignment)
	} else {
		assignment, err = GetAssignmentFac(ctx, user.Id)
		fmt.Println(assignment)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, assignment)

	}
}
