package users

type UsersService struct {
}

func (s *UsersService) Find(params UsersServiceFindParams) ([]User, error) {
	return make([]User, 0), nil
}

func (s *UsersService) FindById(id string) (*User, error) {
	return nil, nil
}

func (s *UsersService) Create(user UserCreate) (User, error) {
	return User{}, nil
}

func (s *UsersService) Update(user UserUpdate) (User, error) {
	return User{}, nil
}

func (s *UsersService) PartialUpdate(id string, user UserPartialUpdate) (User, error) {
	return User{}, nil
}

func (s *UsersService) Delete(id string) error {
	return nil
}
