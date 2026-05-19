package repository

import (
	"github.com/Bannawat101/market-aggregator/internal/model"
	"gorm.io/gorm"
)

type CryptoRepository struct {
	db *gorm.DB
}

func NewCryptoRepository(db *gorm.DB) *CryptoRepository {
	return &CryptoRepository{
		db: db,
	}
}

func (r *CryptoRepository) SavePrice(price *model.CryptoPrice) error {
	return r.db.Create(price).Error
}

func (r *CryptoRepository) GetRecentPrices(limit int) ([]model.CryptoPrice, error) {
	var prices []model.CryptoPrice

	err := r.db.Order("created_at desc").Limit(limit).Find(&prices).Error
	if err != nil {
		return nil, err
	}

	return prices, nil
}
