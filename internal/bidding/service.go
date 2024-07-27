package bidding

import (
	"errors"
	"time"

	"github.com/Just-A-NoobieDev/auction-go-server/internal/auction"
	"github.com/Just-A-NoobieDev/auction-go-server/pkg/websocket"
	"github.com/google/uuid"
)

type BidService interface {
	PlaceBid(b *BidRequest, auctionId uuid.UUID) error
	GetBidsByID(id string) ([]*Bid, error)
}

type service struct {
	bidRepo BidRepository
	auctionRepo auction.AuctionRepository
}

func NewBidService(bidRepo BidRepository, auctionRepo auction.AuctionRepository) BidService {
	return &service{bidRepo, auctionRepo}
}

func (s *service) PlaceBid(b *BidRequest, auctionId uuid.UUID) error {
	auction, err := s.auctionRepo.GetAuctionByID(auctionId)
	if err != nil {
		return err
	}

	if auction.Status != "active" {
		return errors.New("auction is not active")
	}

	if auction.OwnerID == uuid.MustParse(b.UserID) {
		return errors.New("owner cannot bid on their own auction")
	}

	// validate user id if it is a participant
	// isParticipant := false
	// for _, participant := range auction.Participants {
	// 	if participant == uuid.MustParse(b.UserID) {
	// 		isParticipant = true
	// 		break
	// 	}
	// }

	// if !isParticipant {
	// 	return errors.New("user is not a participant")
	// }

	if b.Amount < auction.StartPrice {
		return errors.New("bid amount is less than starting price")
	}

	if b.Amount <= auction.CurrentPrice {
		return errors.New("bid amount is less than or equal to current price")
	}

	newID := uuid.New()

	bid := &Bid{
		ID: newID,
		Amount: b.Amount,
		UserID: uuid.MustParse(b.UserID),
		AuctionID: auctionId,
	}

	err = s.bidRepo.CreateBid(bid)

	if err != nil {
		return err
	}
	
	// Broadcast new bid to all participants
	for _, participant := range auction.Participants {
		websocket.Broadcast <- websocket.Message{
			Type:    "new_bid",
			Content: "A new bid has been placed.",
			AuctionID: auction.ID.String(),
			UserID: participant.String(),
		}
	}

	// Broadcast new bid to owner
	websocket.Broadcast <- websocket.Message{
		Type:    "new_bid",
		Content: "A new bid has been placed.",
		AuctionID: auction.ID.String(),
		UserID: auction.OwnerID.String(),
	}

	auction.CurrentPrice = b.Amount
	auction.UpdatedAt = time.Now()

	return s.auctionRepo.UpdateAuction(auction)
}

func (s *service) GetBidsByID(id string) ([]*Bid, error) {
	auctionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	auction, err := s.auctionRepo.GetAuctionByID(auctionID)
	if err != nil {
		return nil, err
	}
	
	return s.bidRepo.GetBidsByID(auction.ID)
}
