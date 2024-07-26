package user

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(username string) (*User, error)
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