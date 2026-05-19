package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Bannawat101/market-aggregator/internal/model"
	"github.com/Bannawat101/market-aggregator/internal/repository"
)

type BinanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type PriceWorker struct {
	repo  *repository.CryptoRepository
	coins []string
}

func NewPriceWorker(repo *repository.CryptoRepository, coins []string) *PriceWorker {
	return &PriceWorker{
		repo:  repo,
		coins: coins,
	}
}

func (w *PriceWorker) Start() {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for range ticker.C {
			var wg sync.WaitGroup
			for _, coin := range w.coins {
				wg.Add(1)

				go w.fetchAndSave(&wg, coin)
			}
			wg.Wait()
			fmt.Println("✅ [Worker] Data retrieved and saved to the latest updated database!")
		}
	}()
}

func (w *PriceWorker) fetchAndSave(wg *sync.WaitGroup, symbol string) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v\n", symbol, err)
		return
	}
	defer resp.Body.Close()

	var binanceResp BinanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&binanceResp); err != nil {
		log.Printf("Error decoding %s: %v\n", symbol, err)
		return
	}

	priceFloat, err := strconv.ParseFloat(binanceResp.Price, 64)
	if err != nil {
		log.Printf("Error parsing price for %s: %v\n", symbol, err)
		return
	}

	newPrice := &model.CryptoPrice{
		Symbol: binanceResp.Symbol,
		Price:  priceFloat,
	}

	err = w.repo.SavePrice(newPrice)
	if err != nil {
		log.Printf("Error saving %s to DB: %v\n", symbol, err)
	}
}
