package auction

import (
	"errors"

	"github.com/Just-A-NoobieDev/auction-go-server/internal/user"
	"github.com/google/uuid"
)

type AuctionFull struct {
	AuctionData Auction `json:"auction"`
	Owner user.User `json:"owner"`
	Participants []user.User `json:"participants"`
}

type AuctionService interface {
	CreateAuction(a *Auction) error
	GetAuctionByID(id uuid.UUID) (*Auction, *user.User, []user.User, error)
	UpdateAuction(a *Auction) error
	DeleteAuction(id uuid.UUID) error
	ListAuctions(offset int, limit int, sort string, sortDirection string, query string) ([]*AuctionFull, error)
	CountAuctions(query string) (int, error)
	JoinAuction(auctionID uuid.UUID, userID uuid.UUID) error
	LeaveAuction(auctionID uuid.UUID, userID uuid.UUID) error
}

type service struct {
	auctionRepo AuctionRepository
	userRepo user.UserRepository
}

func NewAuctionService(auctionRepo AuctionRepository, userRepo user.UserRepository) AuctionService {
	return &service{
		auctionRepo: auctionRepo,
		userRepo: userRepo,
	}
}

func (s *service) CreateAuction(a *Auction) error {
	return s.auctionRepo.CreateAuction(a)
}

func (s *service) GetAuctionByID(id uuid.UUID) (*Auction, *user.User, []user.User, error) {
	auction, err := s.auctionRepo.GetAuctionByID(id)
	if err != nil {
		return nil, nil, nil, err
	}

	owner, err := s.userRepo.GetUserByID(auction.OwnerID)
	if err != nil {
		return nil, nil, nil, err
	}

	participants := make([]user.User, len(auction.Participants))

	for i, participant := range auction.Participants {
		user, err := s.userRepo.GetUserByID(participant)
		if err != nil {
			return nil, nil, nil, err
		}

		participants[i] = *user
	}

	return auction, owner, participants, nil
}

func (s *service) UpdateAuction(a *Auction) error {
	return s.auctionRepo.UpdateAuction(a)
}

func (s *service) DeleteAuction(id uuid.UUID) error {
	return s.auctionRepo.DeleteAuction(id)
}


func (s *service) ListAuctions(offset int, limit int, sort string, sortDirection string, query string) ([]*AuctionFull, error) {
	auctions, err := s.auctionRepo.ListAuctions(offset, limit, sort, sortDirection, query)
	if err != nil {
		return nil, err
	}

	auctionsFull := make([]*AuctionFull, len(auctions))

	for i, auction := range auctions {
		owner, err := s.userRepo.GetUserByID(auction.OwnerID)
		if err != nil {
			return nil, err
		}

		participants := make([]user.User, len(auction.Participants))

		for j, participant := range auction.Participants {
			user, err := s.userRepo.GetUserByID(participant)
			if err != nil {
				return nil, err
			}

			participants[j] = *user
		}

		auctionsFull[i] = &AuctionFull{
			AuctionData: *auction,
			Owner: *owner,
			Participants: participants,
		}
	}

	return auctionsFull, nil
}

func (s *service) CountAuctions(query string) (int, error) {
	return s.auctionRepo.CountAuctions(query)
}

func (s *service) JoinAuction(auctionID uuid.UUID, userID uuid.UUID) error {
	auction, err := s.auctionRepo.GetAuctionByID(auctionID)
	if err != nil {
		return err
	}

	if auction.OwnerID == userID {
		return errors.New("you can't join your own auction")
	}

	for _, participant := range auction.Participants {
		if participant == userID {
			return errors.New("you already joined this auction")
		}
	}

	auction.Participants = append(auction.Participants, userID)
	return s.auctionRepo.UpdateAuction(auction)
}

func (s *service) LeaveAuction(auctionID uuid.UUID, userID uuid.UUID) error {
	auction, err := s.auctionRepo.GetAuctionByID(auctionID)
	if err != nil {
		return err
	}

	if auction.OwnerID == userID {
		return errors.New("you can't leave your own auction")
	}

	for i, participant := range auction.Participants {
		if participant == userID {
			auction.Participants = append(auction.Participants[:i], auction.Participants[i+1:]...)
			return s.auctionRepo.UpdateAuction(auction)
		}
	}

	return errors.New("you are not a participant of this auction")
}