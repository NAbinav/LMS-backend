package modules

import (
	"dbms/helper"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Course_id int    `json:"course_id"`
	Title     string `json:"title"`
	Order_num int    `json:"order_num"`
	Content   string `json:"content"`
	Link      string `json:"link"`
}

func CreateModuleHandler(c *gin.Context) {
	var module_input Module
	ctx := c.Request.Context()
	c.ShouldBindJSON(&module_input)
	user, err := helper.WhoamI(c)
	if err != nil || user.Role != "instructor" {
		c.JSON(401, "Unauthorised Access")
		return
	}
	err = CreateModule(ctx, module_input.Course_id, module_input.Order_num, module_input.Title, module_input.Content, module_input.Link)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, "Created Successfully")
}
