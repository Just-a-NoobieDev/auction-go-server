package auction

import (
	"net/http"
	"strconv"

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

	newId := uuid.New()

	auction := &Auction{
		ID: newId,
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
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	sort := c.DefaultQuery("sort", "created_at")
	sortDirection := c.DefaultQuery("sort_direction", "asc")
	query := c.DefaultQuery("query", "")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size value"})
		return
	}

	offset := (page - 1) * pageSize

	auctions, err := h.service.ListAuctions(offset, pageSize, sort, sortDirection, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalItems, err := h.service.CountAuctions(query) // Add a method to count total auctions
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (totalItems + pageSize - 1) / pageSize
	remainingItems := totalItems - (page * pageSize)
	if remainingItems < 0 {
		remainingItems = 0
	}
	remainingPages := totalPages - page
	if remainingPages < 0 {
		remainingPages = 0
	}

	response := AuctionListResponse{
		Items: auctions,
		Meta: Meta{
			Page:           page,
			PageSize:       pageSize,
			TotalItems:     totalItems,
			TotalPages:     totalPages,
			RemainingItems: remainingItems,
			RemainingPages: remainingPages,
			Sort:           sort,
			SortDirection:  sortDirection,
			Query:          query,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuctionHandler) JoinAuction(c *gin.Context) {
	auctionID := c.Param("id")

	userID := uuid.MustParse(c.GetString("userID"))

	err := h.service.JoinAuction(uuid.MustParse(auctionID), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *AuctionHandler) LeaveAuction(c *gin.Context) {
	auctionID := c.Param("id")
	id, err := uuid.Parse(auctionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uuid.MustParse(c.GetString("userID"))

	err = h.service.LeaveAuction(id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}