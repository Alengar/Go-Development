package user

type UserRepositoryImpl struct {
	users  map[int]*User
	lastID int
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		users:  make(map[int]*User),
		lastID: 0,
	}
}

func (r *UserRepositoryImpl) GetAllUsers() ([]*User, error) {
	users := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserByID(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(user *User) (*User, error) {
	r.lastID++
	user.ID = r.lastID
	r.users[user.ID] = user
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *User) (*User, error) {
	_, exists := r.users[user.ID]
	if !exists {
		return nil, ErrUserNotFound
	}
	r.users[user.ID] = user
	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int) error {
	_, exists := r.users[id]
	if !exists {
		return ErrUserNotFound
	}
	delete(r.users, id)
	return nil
}
