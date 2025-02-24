package stocks

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

type StockQuote struct {
    Symbol  string  `json:"symbol"`
    Price   float64 `json:"price"`
    ATH     float64 `json:"ath"`
    Close   float64 `json:"close"`
}

func GetStockQuote(symbol string) (StockQuote, error) {
    apiKey := os.Getenv("BR_STOCKS_API_KEY") // Exemplo com BRAPI ou HG Brasil
    if apiKey == "" {
        return StockQuote{}, fmt.Errorf("API Key não encontrada")
    }

    url := fmt.Sprintf("https://brapi.dev/api/quote/%s?apikey=%s", symbol, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return StockQuote{}, err
    }
    defer resp.Body.Close()

    var quote StockQuote
    if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
        return StockQuote{}, err
    }

    return quote, nil
}
