package users

type User struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Age   int      `json:"age"`
	Role  UserRole `json:"role"`
}

type UserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

type UserUpdate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

type UserPartialUpdate struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Age   *int    `json:"age"`
	Role  *string `json:"role"`
}
