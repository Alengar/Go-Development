package user

type UserUseCase interface {
	GetAllUsers() ([]*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}

type userUseCase struct {
	userRepo UserRepository
}

func NewUserUseCase(userRepo UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) GetAllUsers() ([]*User, error) {
	return uc.userRepo.GetAllUsers()
}

func (uc *userUseCase) GetUserByID(id int) (*User, error) {
	return uc.userRepo.GetUserByID(id)
}

func (uc *userUseCase) CreateUser(user *User) (*User, error) {
	return uc.userRepo.CreateUser(user)
}

func (uc *userUseCase) UpdateUser(user *User) (*User, error) {
	return uc.userRepo.UpdateUser(user)
}

func (uc *userUseCase) DeleteUser(id int) error {
	return uc.userRepo.DeleteUser(id)
}
