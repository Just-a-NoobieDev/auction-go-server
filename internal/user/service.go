package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(createUserRequest CreateUserRequest) error
	AuthenticateUser(email, password string) (*User, error)
	GetUser(userID string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(user *UpdateUserRequest, userID string) error
	DeleteUser(userID string) error
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

	newId := uuid.New()


	user := &User{
		ID:           newId,
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
	return s.userRepo.GetUserByID(uuid.MustParse(userID))
}


// TODO: Implement pagination and search
func (s *service) GetAllUsers() ([]User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *service) UpdateUser(user *UpdateUserRequest, userID string) error {
	err := uuid.Validate(userID)

	if err != nil {
		return err
	}

	_, err = s.userRepo.GetUserByID(uuid.MustParse(userID))
	if err != nil {
		return err
	}

	return s.userRepo.UpdateUser(user, userID)
}

func (s *service) DeleteUser(userID string) error {
	err := uuid.Validate(userID)

	if err != nil {
		return err
	}

	_, err = s.userRepo.GetUserByID(uuid.MustParse(userID))
	if err != nil {
		return err
	}

	return s.userRepo.DeleteUser(userID)
}