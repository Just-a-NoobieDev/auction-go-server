package auction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuctionHandler struct {
	service AuctionService
}

func NewAuctionHandler(service AuctionService) *AuctionHandler {
	return &AuctionHandler{service}
}

func (h *AuctionHandler) CreateAuction(c *gin.Context) {
	var createAuctionRequest CreateAuctionRequest
	if err := c.ShouldBindJSON(&createAuctionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	auction := &Auction{
		Title: 	 createAuctionRequest.Title,
		Description: createAuctionRequest.Description,
		StartPrice: createAuctionRequest.StartPrice,
		CurrentPrice: createAuctionRequest.StartPrice,
		StartDate: createAuctionRequest.StartDate,
		EndDate: createAuctionRequest.EndDate,
		OwnerID: uuid.MustParse(c.GetString("userID")),
		MinIncrement: createAuctionRequest.MinIncrement,
		Status: "upcoming",
	}

	err := h.service.CreateAuction(auction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, auction)
}

func (h *AuctionHandler) GetAuctionByID(c *gin.Context) {
	auctionID := c.Param("id")
	id, err := uuid.Parse(auctionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auction, owner, participants, err := h.service.GetAuctionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"auction": auction,
		"owner": owner,
		"participants": participants,
	})
}

func (h *AuctionHandler) UpdateAuction(c *gin.Context) {
	auctionID := c.Param("id")
	id, err := uuid.Parse(auctionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updateAuctionRequest UpdateAuctionRequest
	if err := c.ShouldBindJSON(&updateAuctionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auction, _, _, err := h.service.GetAuctionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	auction.Title = updateAuctionRequest.Title
	auction.Description = updateAuctionRequest.Description
	auction.StartPrice = updateAuctionRequest.StartPrice
	auction.CurrentPrice = updateAuctionRequest.CurrentPrice
	auction.StartDate = updateAuctionRequest.StartDate
	auction.EndDate = updateAuctionRequest.EndDate
	auction.MinIncrement = updateAuctionRequest.MinIncrement
	auction.Status = updateAuctionRequest.Status

	err = h.service.UpdateAuction(auction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, auction)
}

func (h *AuctionHandler) DeleteAuction(c *gin.Context) {
	auctionID := c.Param("id")
	id, err := uuid.Parse(auctionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.DeleteAuction(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *AuctionHandler) ListAuctions(c *gin.Context) {
	auctions, err := h.service.ListAuctions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auctions": auctions})
}

// TODO: Implement SearchAuctions and Pagination
// TODO: Resolve the owner_id to user details and participants to user details