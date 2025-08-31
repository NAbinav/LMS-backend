package handler

import (
	"dbms/db"
	"dbms/helper"
	"dbms/jwt"
	"dbms/user"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Courses struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	InstructorID int       `json:"instructor_id"`
	Credits      int       `json:"credits"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func RegisterUser(c *gin.Context) {
	var User User
	if err := c.ShouldBindJSON(&User); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	ctx := c.Request.Context()
	if err := user.RegisterUser(ctx, User.Name, User.Email, User.Password, User.Role); err != nil {
		fmt.Print("Error registering user:", err)
		c.JSON(500, gin.H{"error": "Failed to register User"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {
	var User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&User); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	ctx := c.Request.Context()
	id, role, err := user.LoginUser(ctx, User.Email, User.Password) // dv change
	if err != nil {
		fmt.Println("Error logging in user:", err)
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}
	token, err := jwt.Create_JWT(User.Email)
	if err != nil {
		c.JSON(401, "unauthorised")
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "Login successful", "id": id, "role": role}) //dv change
}

func Getuser(c *gin.Context) {
	user, err := helper.WhoamI(c)
	if err != nil {
		c.JSON(401, err)
		return
	}
	c.JSON(200, user)
	return
}

func ListUserFromRole(c *gin.Context) {
	ctx := c.Request.Context()
	role := c.Query("role")
	if role == "" {
		c.JSON(400, gin.H{"error": "Role is required"})
		return
	}

	query := "SELECT * FROM users WHERE role = $1"
	rows, err := db.DB.Query(ctx, query, role)
	if err != nil {
		fmt.Println("Error fetching users by role:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
	}

	for rows.Next() {
		var user struct {
			Id        int       `json:"id"`
			Name      string    `json:"name"`
			Email     string    `json:"email"`
			Password  string    `json:"password"`
			Role      string    `json:"role"`
			CreatedAt time.Time `json:"created_at"`
		}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
			fmt.Println("Error scanning user row:", err)
			c.JSON(500, gin.H{"error": "Failed to scan user row"})
			return
		}
		users = append(users, user)
	}

	c.JSON(200, users)
}

func DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	query := "DELETE FROM users WHERE id = $1"
	result, err := db.DB.Exec(ctx, query, id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
