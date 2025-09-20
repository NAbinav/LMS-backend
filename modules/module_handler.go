package modules

import (
	"dbms/helper"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Module struct {
	Course_id int    `json:"course_id"`
	Title     string `json:"title"`
	Order_num int    `json:"order_num"`
	Content   string `json:"content"`
	Link      string `json:"link"`
}

type Module struct {
	Course_id int    json:"course_id"
	Title     string json:"title"
	Content   string json:"content"
	Link      string json:"link"
}

func CreateModuleHandler(c *gin.Context) {
	var module_input Module
	ctx := c.Request.Context()
	c.ShouldBindJSON(&module_input)
	user, err := helper.WhoamI(c)
	fmt.Println(user, err)
	if err != nil || user.Role != "instructor" || !helper.CheckValidFaculty(ctx, user.Id, module_input.Course_id) {
		c.JSON(401, "Unauthorised Access")
		return
	}
	err = CreateModule(ctx, module_input.Course_id, module_input.Title, module_input.Content, module_input.Link)
	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, "Created Successfully")
}

func GetModulesHandler(c *gin.Context) {
	ctx := c.Request.Context()
	mod := c.Query("m_id")
	c_id := c.Query("c_id")
	user, err := helper.WhoamI(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
	}
	if !helper.CheckIfEnrolled(ctx, user.Id, c_id) {
		c.JSON(400, "Cant Access without enrolling ")
		return
	}
	module, err := GetModules(ctx, c_id, mod)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
		return
	}
	c.JSON(200, module)

}
func GetAllModulesHandler(c *gin.Context) {

	ctx := c.Request.Context()
	c_id := c.Query("c_id")
	user, err := helper.WhoamI(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
	}
	if !helper.CheckIfEnrolled(ctx, user.Id, c_id) {
		c.JSON(400, "Cant Access without enrolling ")
		return
	}
	module, err := GetAllModules(ctx, c_id)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, err)
		return
	}
	c.JSON(200, module)
}
