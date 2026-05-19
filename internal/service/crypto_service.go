package service

import (
	"github.com/Bannawat101/market-aggregator/internal/model"
	"github.com/Bannawat101/market-aggregator/internal/repository"
)

type CryptoService struct {
	repo *repository.CryptoRepository
}

func NewCryptoService(repo *repository.CryptoRepository) *CryptoService {
	return &CryptoService{
		repo: repo,
	}
}

func (s *CryptoService) GetLastestPrice(limit int) ([]model.CryptoPrice, error) {
	if limit > 100 {
		limit = 100
	}

	return s.repo.GetRecentPrices(limit)
}
