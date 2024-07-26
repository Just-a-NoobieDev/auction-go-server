package userprofile

import "gorm.io/gorm"

type UserProfileRepository interface {
	CreateUserProfile(userProfile *UserProfile) error
	GetUserProfileByUserID(userID string) (*UserProfile, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
	return &repository{db}
}

func (r *repository) CreateUserProfile(userProfile *UserProfile) error {
	return r.db.Create(userProfile).Error
}

func (r *repository) GetUserProfileByUserID(userID string) (*UserProfile, error) {
	var userProfile UserProfile
	err := r.db.Where("user_id = ?", userID).First(&userProfile).Error
	return &userProfile, err
}

