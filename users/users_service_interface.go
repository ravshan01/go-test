package users

type IUsersService interface {
	find(params UsersServiceFindParams) (*FindUsersResponse, error)
	findById(id string) (*User, error)
	create(user UserCreate) (*User, error)
	update(user UserUpdate) (*User, error)
	partialUpdate(id string, user UserPartialUpdate) (*User, error)
	delete(id string) error
}

type UsersServiceFindParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
