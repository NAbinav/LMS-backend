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
type Grading struct {
	Grade int `json:"grade"`
}

func NewSubmissionHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var Submission NewSubmission
	err := c.ShouldBindJSON(&Submission)
	if err != nil {
		c.AbortWithError(400, err)
	}
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
		c.AbortWithError(400, errors.New("not authorised"))
		return
	}
	val, err := AllSubmissions(ctx, a_id)
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.JSON(200, val)
}

func GradeAssignment(c *gin.Context) {
	ctx := c.Request.Context()
	var Grade Grading
	err := c.ShouldBindJSON(&Grade)
	if err != nil {
		c.AbortWithError(400, err)
	}
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(200, err)
		return
	}
	fmt.Println(user, Grade.Grade)
	isAssigned := CheckIfAssigned(ctx, user.Id, Grade.Grade)
	if !isAssigned {
		c.JSON(401, "Not Assigned")
		return
	}
	err = GradeAssn_db(ctx, user.Id, Grade.Grade)
	if err != nil {
		c.JSON(400, err)
		return
	}
}

func CheckSubmissionHandler(c *gin.Context) {
	var req struct {
		UserID       int `json:"user_id"`
		AssignmentID int `json:"assignment_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	submitted, err := HasSubmitted(c.Request.Context(), req.UserID, req.AssignmentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	c.JSON(200, gin.H{"submitted": submitted})
}
