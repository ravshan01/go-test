package users

type UsersService struct {
}

func (s *UsersService) find(params UsersServiceFindParams) ([]User, error) {
	return nil, nil
}

func (s *UsersService) findById(id string) (User, error) {
	return User{}, nil
}

func (s *UsersService) create(user UserCreate) (User, error) {
	return User{}, nil
}

func (s *UsersService) update(user UserUpdate) (User, error) {
	return User{}, nil
}

func (s *UsersService) partialUpdate(id string, user UserPartialUpdate) (User, error) {
	return User{}, nil
}

func (s *UsersService) delete(id string) error {
	return nil
}
