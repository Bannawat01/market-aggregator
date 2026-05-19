package main

import (
	"log"
	"os"
	"strings"

	"github.com/Bannawat101/market-aggregator/internal/handler"
	"github.com/Bannawat101/market-aggregator/internal/model"
	"github.com/Bannawat101/market-aggregator/internal/repository"
	"github.com/Bannawat101/market-aggregator/internal/service"
	"github.com/Bannawat101/market-aggregator/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, relying on system environment variables")
	}

	db, err := gorm.Open(sqlite.Open("market.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database")
	}

	db.AutoMigrate(&model.CryptoPrice{})
	log.Println("✅ Database connected and migrated")

	cryptoRepo := repository.NewCryptoRepository(db)
	cryptoService := service.NewCryptoService(cryptoRepo)
	cryptoHandler := handler.NewCryptoHandler(cryptoService)

	coinsEnv := os.Getenv("TARGET_COINS")
	coins := strings.Split(coinsEnv, ",")
	priceWorker := worker.NewPriceWorker(cryptoRepo, coins)

	priceWorker.Start()

	r := gin.Default()

	r.GET("/api/v1/prices", cryptoHandler.GetPrices)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server is running on port %s", port)
	r.Run(":" + port)
}
