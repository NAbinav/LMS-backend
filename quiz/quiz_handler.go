package quiz

import (
	// "context"
	"dbms/helper"
	// "dbms/quiz"
	"fmt"

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
	// user, err := helper.WhoamI(c)
	// if err != nil {
	// 	c.JSON(401, "Not Authenticated to create quiz")
	// }
	// if !helper.CheckRole(c, "instructor") {
	// 	c.JSON(401, "Not Authenticated to create quiz")
	// 	return
	// }
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
	if user.Role == "student" {
		quiz := AllQuizEnrolled(ctx, user.Id)
		c.JSON(200, quiz)
	} else {
		quiz, err := GetQuizzesByTeacher(ctx, user.Id)
		if err != nil {
			fmt.Println(err)
			c.AbortWithError(400, err)
			return
		}
		c.JSON(200, quiz)
	}
}

// func GetQuizHandler(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	q_id := c.Query("id")
// 	user, err := helper.WhoamI(c)
// 	if err != nil {
// 		c.AbortWithError(200, err)
// 	}
// }
