package crypto

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"

	"github.com/theus-ortiz/invest-lib-go/config"
)

type CryptoQuote struct {
	Symbol      string    `json:"symbol"`
	Price       float64   `json:"current_price"`
	ATH         float64   `json:"ath"`
	Close       float64   `json:"close"`
	MaxSupply   float64   `json:"max_supply"`
	LastUpdated string `json:"last_updated"`
}

func GetCryptoQuote(symbol string) (CryptoQuote, error) {
	config.LoadEnv()

	apiKey := config.GetAPIKey()
	fmt.Println("API Key encontrada:", apiKey)

	// Corrigido o erro no URL da API
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s", symbol)
	resp, err := http.Get(url)
	if err != nil {
		return CryptoQuote{}, err
	}
	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("erro:",err)
	// 	panic(1)
	// }

	// fmt.Println(string(body))

	// Ajustado para refletir que a resposta é uma lista de objetos
	var quotes []CryptoQuote
	if err := json.NewDecoder(resp.Body).Decode(&quotes); err != nil {
		return CryptoQuote{}, err
	}

	// Verifica se a lista de quotes contém o símbolo desejado
	if len(quotes) == 0 {
		return CryptoQuote{}, fmt.Errorf("quote não encontrada para o símbolo %s", symbol)
	}

	// Retorna o primeiro item da lista (se houver)
	return quotes[0], nil
}
