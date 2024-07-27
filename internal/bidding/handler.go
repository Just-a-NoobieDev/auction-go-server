package bidding

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BidHandler struct {
	bidService BidService
}

func NewBidHandler(bidService BidService) *BidHandler {
	return &BidHandler{bidService}
}

func (h *BidHandler) PlaceBid(c *gin.Context) {
	var req BidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	auctionID, err := uuid.Parse(c.Param("id"))	

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("userID")

	req.UserID = userID

	err = h.bidService.PlaceBid(&req, auctionID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "bid placed successfully"})
}

func (h *BidHandler) GetBidsByID(c *gin.Context) {
	auctionID := c.Param("id")

	bids, err := h.bidService.GetBidsByID(auctionID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bids)
}