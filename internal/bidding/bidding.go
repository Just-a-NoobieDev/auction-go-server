package bidding

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key" db:"id"`
	Amount    float64   `json:"amount" gorm:"not null" db:"amount"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null" db:"user_id"`
	AuctionID uuid.UUID `json:"auction_id" gorm:"type:uuid;not null" db:"auction_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" db:"created_at"`
}

type BidRequest struct {
	Amount float64 `json:"amount"`
	UserID string `json:"user_id,omitempty"`
}