package userprofile

import "github.com/google/uuid"

type UserProfile struct {
	UserID            uuid.UUID `json:"user_id" db:"user_id"`
	Bio               string    `json:"bio" db:"bio"`
	ProfilePictureURL string    `json:"profile_picture_url" db:"profile_picture_url"`
	Address           string    `json:"address" db:"address"`
	City              string    `json:"city" db:"city"`
	State             string    `json:"state" db:"state"`
	PostalCode        string    `json:"postal_code" db:"postal_code"`
	Country           string    `json:"country" db:"country"`
}