package schema

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Courses struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	InstructorID int       `json:"instructor_id"`
	Credits      int       `json:"credits"`
	CreatedAt    time.Time `json:"created_at"`
}
type Enrollment struct {
	UserId     int       `json:"user_id"`
	CourseID   int       `json:"course_id"`
	EnrollDate time.Time `json:"enrollment_date"`
	Status     string    `json:"status"`
}
