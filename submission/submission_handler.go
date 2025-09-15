package submission

import (
	// "context"
	"dbms/helper"
	"errors"
	"fmt"

	// "dbms/schema"

	"github.com/gin-gonic/gin"
)

type NewSubmission struct {
	AssignmentId   int    `json:"assignment_id"`
	SubmissionText string `json:"submission_text"`
}

func NewSubmissionHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var Submission NewSubmission
	err := c.ShouldBindJSON(&Submission)
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(200, err)
		return
	}
	fmt.Println(user, Submission.AssignmentId)
	isAssigned := CheckIfAssigned(ctx, user.Id, Submission.AssignmentId)
	if !isAssigned {
		c.JSON(401, "Not Assigned")
		return
	}
	err = AssignmentSubmit(ctx, user.Id, Submission.AssignmentId, Submission.SubmissionText)
	if err != nil {
		c.JSON(400, err)
		return
	}
}

func GetAllSubmissions(c *gin.Context) {
	ctx := c.Request.Context()
	a_id := c.Query("a_id")
	c_id, err := CourseidOfAssignment(ctx, a_id)
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.JSON(200, c_id)
	user, err := helper.WhoamI(c)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	if user.Role != "instructor" || !helper.CheckValidFaculty(ctx, user.Id, c_id) {
		c.AbortWithError(400, errors.New("Not authorised"))
		return
	}
	val, err := AllSubmissions(ctx, a_id)
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.JSON(200, val)
}
