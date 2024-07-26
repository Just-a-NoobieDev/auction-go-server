package auction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuctionRepository interface {
	CreateAuction(a *Auction) error
	GetAuctionByID(id uuid.UUID) (*Auction, error)
	UpdateAuction(a *Auction) error
	DeleteAuction(id uuid.UUID) error
	ListAuctions() ([]*Auction, error)
}

type repository struct {
	db *gorm.DB
}

func NewAuctionRepository(db *gorm.DB) AuctionRepository {
	return &repository{db}
}

func (r *repository) CreateAuction(a *Auction) error {
	return r.db.Create(a).Error
}

func (r *repository) GetAuctionByID(id uuid.UUID) (*Auction, error) {
	var auction Auction
	err := r.db.Where("id = ?", id).First(&auction).Error
	return &auction, err
}

func (r *repository) UpdateAuction(a *Auction) error {
	return r.db.Save(a).Error
}

func (r *repository) DeleteAuction(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&Auction{}).Error
}

func (r *repository) ListAuctions() ([]*Auction, error) {
	var auctions []*Auction
	err := r.db.Find(&auctions).Error
	if err != nil {
		return nil, err
	}

	return auctions, err
}