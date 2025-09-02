package quiz

import (
	"dbms/helper"
	// "dbms/schema"
	"github.com/gin-gonic/gin"
)

func CreateQuizHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var quiz_input struct {
		Course_id    int    `json:"course_id"`
		Title        string `json:"title"`
		Max_attempts int    `json:"max_attempts"`
		Time_limit   int    `json:"time_limit"`
	}
	if !helper.CheckRole(c, "instructor") {
		c.JSON(401, "Not Authenticated to create quiz")
		return
	}
	err := c.ShouldBindJSON(&quiz_input)
	if err != nil {
		c.JSON(400, err)
		return
	}
	err = CreateQuiz(ctx, quiz_input.Course_id, quiz_input.Title, quiz_input.Max_attempts, quiz_input.Time_limit)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(201, "Inserted successfully")
}

// NOTE: Gets all the quiz that is available to the user

func GetAllQuizHandler(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(400, "Not available")
		return
	}
	quiz := AllQuizEnrolled(ctx, user.Id)
	c.JSON(200, quiz)

}
