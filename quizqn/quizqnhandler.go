package quizqn

import (
	"dbms/schema"
	// "fmt"

	"github.com/gin-gonic/gin"
)

func AddQns(c *gin.Context) {
	var arr []schema.QuizQuestion
	c.ShouldBindJSON(&arr)
	AddQn_db(arr)
}
