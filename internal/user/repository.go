package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *UpdateUserRequest, userID string) error
	GetAllUsers() ([]User, error)
	DeleteUser(userID string) error
	GetUserByID(userID uuid.UUID) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db}
}

func (r *repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *repository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *repository) UpdateUser(user *UpdateUserRequest, userID string) error {
	return r.db.Model(&User{}).Where("id = ?", userID).Updates(User{
		Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Phone: user.Phone,
		Role: user.Role,
		Address: user.Address,
		ProfilePicture: user.ProfilePicture,
	}).Error
}

func (r *repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *repository) DeleteUser(userID string) error {
	return r.db.Where("id = ?", userID).Delete(&User{}).Error
}

func (r *repository) GetUserByID(userID uuid.UUID) (*User, error) {
	var user User
	err := r.db.Where("id = ?", userID).First(&user).Error
	return &user, err
}