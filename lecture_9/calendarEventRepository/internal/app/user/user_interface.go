package user

type UserRepository interface {
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}
