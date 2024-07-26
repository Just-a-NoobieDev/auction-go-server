package auction

import (
	"time"

	"github.com/google/uuid"
)

type Auction struct {
	ID           uuid.UUID   `json:"id" gorm:"type:uuid;primary_key" db:"id"`
	Title        string      `json:"title" gorm:"not null" db:"title"`
	Description  string      `json:"description,omitempty" db:"description"`
	StartPrice   float64     `json:"start_price" gorm:"not null" db:"start_price"`
	CurrentPrice float64     `json:"current_price" gorm:"not null;default:0" db:"current_price"`
	StartDate    time.Time   `json:"start_date" gorm:"not null" db:"start_date"`
	EndDate      time.Time   `json:"end_date" gorm:"not null" db:"end_date"`
	Status       string      `json:"status" gorm:"not null;default:'upcoming'" db:"status"`
	MinIncrement float64     `json:"min_increment" gorm:"not null" db:"min_increment"`
	OwnerID      uuid.UUID   `json:"owner_id" gorm:"type:uuid;not null" db:"owner_id"`
	Participants []uuid.UUID `json:"participants" gorm:"type:uuid[]" db:"participants"`
	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime" db:"updated_at"`
}

type CreateAuctionRequest struct {
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description"`
	StartPrice   float64   `json:"start_price" binding:"required"`
	StartDate    time.Time `json:"start_date" binding:"required"`
	EndDate      time.Time `json:"end_date" binding:"required"`
	MinIncrement float64   `json:"min_increment" binding:"required"`
}

type UpdateAuctionRequest struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StartPrice   float64   `json:"start_price"`
	CurrentPrice float64   `json:"current_price"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	MinIncrement float64   `json:"min_increment"`
	Status       string    `json:"status"`
}

type AuctionListResponse struct {
	Items []*Auction `json:"items"`
	Meta  Meta              `json:"meta"`
}

type Meta struct {
	Page           int    `json:"page"`
	PageSize       int    `json:"page_size"`
	TotalItems     int    `json:"total_items"`
	TotalPages     int    `json:"total_pages"`
	RemainingItems int    `json:"remaining_items"`
	RemainingPages int    `json:"remaining_pages"`
	Sort           string `json:"sort"`
	SortDirection  string `json:"sort_direction"`
	Query          string `json:"query"`
}
