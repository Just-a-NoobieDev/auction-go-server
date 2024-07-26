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
	ListAuctions(offset int, limit int, sort string, sortDirection string, query string) ([]*Auction, error)
	CountAuctions(query string) (int, error)
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

func (r *repository) ListAuctions(offset int, limit int, sort string, sortDirection string, query string) ([]*Auction, error) {
	var auctions []*Auction
	queryBuilder := r.db.Order(sort + " " + sortDirection).Offset(offset).Limit(limit)
	if query != "" {
		queryBuilder = queryBuilder.Where("title ILIKE ?", "%"+query+"%")
	}
	if err := queryBuilder.Find(&auctions).Error; err != nil {
		return nil, err
	}
	return auctions, nil
}

func (r *repository) CountAuctions(query string) (int, error) {
	var count int64
	queryBuilder := r.db.Model(&Auction{})
	if query != "" {
		queryBuilder = queryBuilder.Where("title ILIKE ?", "%"+query+"%")
	}
	if err := queryBuilder.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}