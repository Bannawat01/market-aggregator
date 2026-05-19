package model

import "gorm.io/gorm"

type CryptoPrice struct {
	gorm.Model
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
