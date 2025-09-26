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
	Sid   int `json:"s_id"`
}

func NewSubmissionHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var Submission NewSubmission
	err := c.ShouldBindJSON(&Submission)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		c.JSON(401, "Not Assigned")
		return
	}
	err = AssignmentSubmit(ctx, user.Id, Submission.AssignmentId, Submission.SubmissionText)
	if err != nil {
		fmt.Println(err)
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
		fmt.Print(err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "Authentication error"})
		return
	}

	fmt.Printf("Received Grade: %d, S_id: %d\n", Grade.Grade, Grade.Sid)

	err = GradeAssignment_db(ctx, Grade.Sid, Grade.Grade, user.Id)
	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{"error": "Failed to grade"})
		return
	}

	c.JSON(200, gin.H{"message": "Grade submitted successfully"})
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

func GetUserSubmissionsHandler(c *gin.Context) {
	user, err := helper.WhoamI(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	submissions, err := GetUserSubmissions(c.Request.Context(), user.Id)
	if err != nil {
		c.JSON(500, gin.H{"error": "db error"})
		return
	}

	c.JSON(200, submissions)
}
