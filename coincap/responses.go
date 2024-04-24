package coincap

import "fmt"

type CoinResponse struct {
	Coin      Coin  `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Coin struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUSD      string `json:"marketCapUsd"`
	VolumeUSD24Hr     string `json:"volumeUsd24Hr"`
	PriceUSD          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	VWap24Hr          string `json:"vwap24Hr"`
}

type MarketsList struct {
	Market    []Market `json:"data"`
	Timestamp int64    `json:"timestamp"`
}

type Market struct {
	ExchangeID    string `json:"exchangeId"`
	BaseID        string `json:"baseId"`
	QuoteID       string `json:"quoteId"`
	BaseSymbol    string `json:"baseSymbol"`
	QuoteSymbol   string `json:"quoteSymbol"`
	VolumeUSD24Hr string `json:"volumeUsd24Hr"`
	PriceUSD      string `json:"priceUsd"`
	VolumePercent string `json:"volumePercent"`
}

func (c Coin) Info() string {
	return fmt.Sprintf("ID : %s | RANK : %s | SYMBOL : %s | NAME : %s | PRICE : %s",
		c.ID, c.Rank, c.Symbol, c.Name, c.PriceUSD)
}

func (m Market) MarketsInfo() string{
	return fmt.Sprintf("ExchangeID : %s | BaseID : %s | QuoteID : %s | BaseSymbol : %s | QuoteSymbol : %s | VolumeUSD24Hr : %s | PriceUSD : %s | VolumePercent : %s", 
	m.ExchangeID, m.BaseID, m.QuoteID, m.BaseSymbol, m.QuoteSymbol, m.VolumeUSD24Hr, m.PriceUSD, m.VolumePercent)
}