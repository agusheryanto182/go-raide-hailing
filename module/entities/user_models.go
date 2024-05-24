package entities

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Email    string `db:"email"`
	// CreatedAt time.Time `db:"created_at"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)
