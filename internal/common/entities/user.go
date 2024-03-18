package entities

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	ID       string
	Username string
	Email    string
	Role     Role
}
