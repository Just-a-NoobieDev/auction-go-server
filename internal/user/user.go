package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key" db:"id"`  
	Username     string    `json:"username" gorm:"unique;not null" db:"username"`
	Email        string    `json:"email" gorm:"unique;not null" db:"email"`
	PasswordHash string    `json:"-" gorm:"not null" db:"password_hash"` 
	FirstName    string    `json:"first_name" gorm:"not null" db:"first_name"`
	LastName     string    `json:"last_name" gorm:"not null" db:"last_name"`
	Phone        string    `json:"phone" gorm:"not null" db:"phone"`
	CreatedAt    time.Time `json:"created_at"  db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"  db:"updated_at"`
	LastLogin    *time.Time `json:"last_login,omitempty" db:"last_login"` 
}

type CreateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

