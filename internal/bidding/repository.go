package bidding

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BidRepository interface {
	CreateBid(b *Bid) error
	GetBidsByID(id uuid.UUID) ([]*Bid, error)
	// ListBids(offset int, limit int, sort string, sortDirection string, query string) ([]*Bid, error)
	// CountBids(query string) (int, error)
}

type repository struct {
	db *gorm.DB
}

func NewBidRepository(db *gorm.DB) BidRepository {
	return &repository{db}
}

func (r *repository) CreateBid(b *Bid) error {
	return r.db.Create(b).Error
}

func (r *repository) GetBidsByID(id uuid.UUID) ([]*Bid, error) {
	var bids []*Bid
	if err := r.db.Where("auction_id = ?", id).Find(&bids).Error; err != nil {
		return nil, err
	}
	return bids, nil
}