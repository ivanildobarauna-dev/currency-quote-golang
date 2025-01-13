package entities

import "time"

type CurrencyQuote struct {
	CurrencyPairCodes    *Currency
	CurrencyPairFullName string
	BaseCurrencyCode     string
	QuoteCurrencyCode    string
	QuoteTimestamp       int
	BidPrice             float64
	AskPrice             float64
	QuoteExtractedAt     time.Time
}

func NewCurrencyQuote(currencyPairCodes *Currency, currencyPairName string, baseCurrencyCode string, quoteCurrencyCode string, quoteTimestamp int, bidPrice float64, askPrice float64) *CurrencyQuote {
	currQuote := CurrencyQuote{
		CurrencyPairCodes:    currencyPairCodes,
		CurrencyPairFullName: currencyPairName,
		BaseCurrencyCode:     baseCurrencyCode,
		QuoteCurrencyCode:    quoteCurrencyCode,
		BidPrice:             bidPrice,
		AskPrice:             askPrice,
		QuoteTimestamp:       quoteTimestamp,
		QuoteExtractedAt:     time.Now(),
	}

	return &currQuote
}
