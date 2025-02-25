package crypto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CryptoQuote represents the structured cryptocurrency data
type CryptoQuote struct {
	ID                   string     `json:"id"`
	Symbol               string     `json:"symbol"`
	Name                 string     `json:"name"`
	Image                string     `json:"image"`
	PriceInfo            PriceData  `json:"price_info"`
	MarketInfo           MarketData `json:"market_info"`
	VolumeInfo           VolumeData `json:"volume_info"`
	SupplyInfo           SupplyData `json:"supply_info"`
	ATHInfo              ATHData    `json:"ath_info"`
	ATLInfo              ATLData    `json:"atl_info"`
	LastUpdated          string     `json:"last_updated"`
	FormattedLastUpdated string
}

// Subgroup structures
type PriceData struct {
	CurrentPrice      float64  `json:"current_price"`
	High24h           float64  `json:"high_24h"`
	Low24h            float64  `json:"low_24h"`
	PriceChange24h    float64  `json:"price_change_24h"`
	PriceChangePct24h float64  `json:"price_change_percentage_24h"`
}

type MarketData struct {
	MarketCap             float64 `json:"market_cap"`
	MarketCapRank         int     `json:"market_cap_rank"`
	MarketCapChange24h    float64 `json:"market_cap_change_24h"`
	MarketCapChangePct24h float64 `json:"market_cap_change_percentage_24h"`
}

type VolumeData struct {
	TotalVolume float64 `json:"total_volume"`
}

type SupplyData struct {
	CirculatingSupply float64 `json:"circulating_supply"`
	TotalSupply       float64 `json:"total_supply"`
	MaxSupply         float64 `json:"max_supply"`
}

type ATHData struct {
	ATH          float64 `json:"ath"`
	ATHChangePct float64 `json:"ath_change_percentage"`
	ATHDate      string  `json:"ath_date"`
}

type ATLData struct {
	ATL          float64 `json:"atl"`
	ATLChangePct float64 `json:"atl_change_percentage"`
	ATLDate      string  `json:"atl_date"`
}

// Converts interface{} to float64 with nil check
func toFloat64(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return v
	}
	return 0.0 // Returns 0.0 if value is nil or not a float
}

// Converts interface{} to string with nil check
func toString(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	}
	return "" // Returns an empty string if value is nil
}

// GetCryptoQuote fetches data from the CoinGecko API
func GetCryptoQuote(symbol string) (CryptoQuote, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s", symbol)
	resp, err := http.Get(url)
	if err != nil {
		return CryptoQuote{}, err
	}
	defer resp.Body.Close()

	var rawData []map[string]interface{} // Decoding generic JSON first
	if err := json.NewDecoder(resp.Body).Decode(&rawData); err != nil {
		return CryptoQuote{}, err
	}

	if len(rawData) == 0 {
		return CryptoQuote{}, fmt.Errorf("❌ No data found for %s", symbol)
	}

	data := rawData[0] // Take the first cryptocurrency returned

	// Creating the structure while preventing conversion errors
	quote := CryptoQuote{
		ID:      toString(data["id"]),
		Symbol:  toString(data["symbol"]),
		Name:    toString(data["name"]),
		Image:   toString(data["image"]),
		PriceInfo: PriceData{
			CurrentPrice:       toFloat64(data["current_price"]),
			High24h:            toFloat64(data["high_24h"]),
			Low24h:             toFloat64(data["low_24h"]),
			PriceChange24h:     toFloat64(data["price_change_24h"]),
			PriceChangePct24h:  toFloat64(data["price_change_percentage_24h"]),
		},
		MarketInfo: MarketData{
			MarketCap:             toFloat64(data["market_cap"]),
			MarketCapRank:         int(toFloat64(data["market_cap_rank"])), // Converting float64 to int
			MarketCapChange24h:    toFloat64(data["market_cap_change_24h"]),
			MarketCapChangePct24h: toFloat64(data["market_cap_change_percentage_24h"]),
		},
		VolumeInfo: VolumeData{
			TotalVolume: toFloat64(data["total_volume"]),
		},
		SupplyInfo: SupplyData{
			CirculatingSupply: toFloat64(data["circulating_supply"]),
			TotalSupply:       toFloat64(data["total_supply"]),
			MaxSupply:         toFloat64(data["max_supply"]),
		},
		ATHInfo: ATHData{
			ATH:          toFloat64(data["ath"]),
			ATHChangePct: toFloat64(data["ath_change_percentage"]),
			ATHDate:      toString(data["ath_date"]),
		},
		ATLInfo: ATLData{
			ATL:          toFloat64(data["atl"]),
			ATLChangePct: toFloat64(data["atl_change_percentage"]),
			ATLDate:      toString(data["atl_date"]),
		},
		LastUpdated: toString(data["last_updated"]),
	}

	// Converting "LastUpdated" to a more readable format
	parsedTime, err := time.Parse(time.RFC3339, quote.LastUpdated)
	if err == nil {
		quote.FormattedLastUpdated = parsedTime.Format("02/01/2006 15:04")
	} else {
		quote.FormattedLastUpdated = quote.LastUpdated // Fallback in case of error
	}

	return quote, nil
}