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
	symbol := "bitcoin" // Nome correto no CoinGecko
	quote, _ := crypto.GetCryptoQuote(symbol)

	fmt.Println("💰 Cotação de", quote.Name, "("+quote.Symbol+")")
	fmt.Println("🔹 Preço Atual:", quote.PriceInfo.CurrentPrice)
	fmt.Println("🚀 Máxima Histórica:", quote.ATHInfo.ATH, "em", quote.ATHInfo.ATHDate)
	fmt.Println("📉 Mínima Histórica:", quote.ATLInfo.ATL, "em", quote.ATLInfo.ATLDate)
	fmt.Println("📊 Market Cap:", quote.MarketInfo.MarketCap)
	fmt.Println("🏆 Rank no Mercado:", quote.MarketInfo.MarketCapRank)
	fmt.Println("📈 Variação 24h:", quote.PriceInfo.PriceChange24h, "(", quote.PriceInfo.PriceChangePct24h, "% )")
	fmt.Println("🔄 Última atualização:", quote.FormattedLastUpdated)
}
