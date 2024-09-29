package users

type User struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Age   int      `json:"age"`
	Role  UserRole `json:"role"`
}
