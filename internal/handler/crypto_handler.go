package handler

import (
	"net/http"
	"strconv"

	"github.com/Bannawat101/market-aggregator/internal/service"
	"github.com/gin-gonic/gin"
)

type CryptoHandler struct {
	service *service.CryptoService
}

func NewCryptoHandler(service *service.CryptoService) *CryptoHandler {
	return &CryptoHandler{
		service: service,
	}
}

func (h *CryptoHandler) GetPrices(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	prices, err := h.service.GetLastestPrice(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch prices"})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    prices,
	})

}
