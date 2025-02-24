package main

import (
	"fmt"

	"github.com/theus-ortiz/invest-lib-go/pkg/crypto"
	//"github.com/theus-ortiz/invest-lib-go/pkg/stocks"
)

func main() {
	// // Buscar cotação de ação brasileira (exemplo: PETR4)
	// stock, err := stocks.GetStockQuote("PETR4")
	// if err != nil {
	//     fmt.Println("Erro ao buscar ação:", err)
	// } else {
	//     fmt.Printf("Ação: %s - Preço Atual: %.2f - ATH: %.2f - Fechamento: %.2f\n",
	//         stock.Symbol, stock.Price, stock.ATH, stock.Close)
	// }

	// Buscar cotação de criptomoeda (exemplo: Bitcoin)
	crypto, err := crypto.GetCryptoQuote("bitcoin")
	if err != nil {
		fmt.Println("Erro ao buscar cripto:", err)
	} else {
		fmt.Printf("Cripto: %s - Preço Atual: %.2f - ATH: %.2f - Max Supply: %.f - Last_Update: %s \n",
			crypto.Symbol, crypto.Price, crypto.ATH, crypto.MaxSupply, crypto.LastUpdated)
	}
}
