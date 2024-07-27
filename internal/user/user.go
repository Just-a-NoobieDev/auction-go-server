package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primary_key" db:"id"`
	Username       string     `json:"username" gorm:"unique;not null" db:"username"`
	Email          string     `json:"email" gorm:"unique;not null" db:"email"`
	PasswordHash   string     `json:"-" gorm:"not null" db:"password_hash"`
	FirstName      string     `json:"first_name" gorm:"not null" db:"first_name"`
	LastName       string     `json:"last_name" gorm:"not null" db:"last_name"`
	Phone          string     `json:"phone" gorm:"not null" db:"phone"`
	Role           string     `json:"role,omitempty" gorm:"not null;default:user" db:"role"`
	Address        string     `json:"address,omitempty" db:"address"`
	ProfilePicture string     `json:"profile_picture,omitempty" db:"profile_picture"`
	CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
	LastLogin      *time.Time `json:"last_login,omitempty" db:"last_login"`
}

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

type AuthenticateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Phone          string `json:"phone"`
	Role           string `json:"role"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
}
