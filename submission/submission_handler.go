package submission

import (
	"dbms/helper"

	"github.com/gin-gonic/gin"
)

type NewSubmission struct {
	AssignmentId   int    `json:"assignment_id"`
	SubmissionText string `json:"submission_text"`
}

func NewSubmissionHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var Submission NewSubmission
	err := c.ShouldBindJSON(Submission)
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(200, err)
		return
	}
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
