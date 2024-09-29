package users

type IUsersService interface {
	Find(params UsersServiceFindParams) ([]User, error)
	FindById(id string) (*User, error)
	Create(user UserCreate) (User, error)
	Update(user UserUpdate) (User, error)
	PartialUpdate(id string, user UserPartialUpdate) (User, error)
	Delete(id string) error
}

type UsersServiceFindParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
