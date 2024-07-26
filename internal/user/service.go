package user

import "golang.org/x/crypto/bcrypt"

type UserService interface {
	RegisterUser(createUserRequest CreateUserRequest) error
	AuthenticateUser(email, password string) (*User, error)
	GetUser(userID string) (*User, error)
}

type service struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &service{userRepo}
}

func (s *service) RegisterUser(createUserRequest CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &User{
		Username:     createUserRequest.Username,
		Email:        createUserRequest.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    createUserRequest.FirstName,
		LastName:     createUserRequest.LastName,
		Phone:        createUserRequest.Phone,
	}

	return s.userRepo.CreateUser(user)
}

func (s *service) AuthenticateUser(email, password string) (*User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUser(userID string) (*User, error) {
	return s.userRepo.GetUserByUsername(userID)
}