package submission

import (
	"dbms/helper"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type NewSubmission struct {
	Id             int       `json:"id"`
	AssignmentId   int       `json:"assignment_id"`
	SubmissionText string    `json:"submission_text"`
	SubmittedAt    time.Time `json:"submitted_at"`
}

func NewSubmissionHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var Submission NewSubmission
	err := c.ShouldBindJSON(Submission)
	if err != nil {
		c.JSON(200, err)
	}
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(200, err)
	}
	ass_id := c.Query("a_id")
	int_ass_id, err := strconv.Atoi(ass_id)
	CheckIfAssigned(ctx, user.Id, int_ass_id)
}
