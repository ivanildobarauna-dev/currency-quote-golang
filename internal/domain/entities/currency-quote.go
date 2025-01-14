package entities

import "time"

type CurrencyQuote struct {
	CurrencyPairCodes    *CurrencyPairList
	CurrencyPairFullName string
	BaseCurrencyCode     string
	QuoteCurrencyCode    string
	QuoteTimestamp       int
	BidPrice             float64
	AskPrice             float64
	QuoteExtractedAt     time.Time
}

func NewCurrencyQuote(currencyPairCodes *CurrencyPairList, currencyPairName string, baseCurrencyCode string, quoteCurrencyCode string, quoteTimestamp int, bidPrice float64, askPrice float64) *[]CurrencyQuote {
	CurrenciesQuote := []CurrencyQuote{CurrencyQuote{
		CurrencyPairCodes:    currencyPairCodes,
		CurrencyPairFullName: currencyPairName,
		BaseCurrencyCode:     baseCurrencyCode,
		QuoteCurrencyCode:    quoteCurrencyCode,
		BidPrice:             bidPrice,
		AskPrice:             askPrice,
		QuoteTimestamp:       quoteTimestamp,
		QuoteExtractedAt:     time.Now(),
	}}

	return &CurrenciesQuote
}
