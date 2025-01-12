package entities

import (
	"time"
	"errors"
)


type currency struct {
	code string
}

func NewCurrency(code string) (*currency, error){
	if len(code) > 3{
		return nil, errors.New("currency code should be on max 3 characters")
	}

	return &currency{code: code}, nil
}

type currencyQuote struct { 
	CurrencyPairCodes [2]currency
	CurrencyPairFullName string
	BaseCurrencyCode currency
	QuoteCurrencyCode currency
	QuoteTimestamp int
	BidPrice float64
	AskPrice float64
	QuoteExtractedAt time.Time
}

func NewCurrencyQuote(currency_pair_codes [2]currency, currency_pair_name string, base_currency_code currency, quote_currency_code currency, quote_timestamp int, bid_price float64, ask_price float64) currencyQuote {
	curr:= currencyQuote{
		CurrencyPairCodes: currency_pair_codes,
		CurrencyPairFullName: currency_pair_name,
		BaseCurrencyCode: base_currency_code,
		QuoteCurrencyCode: quote_currency_code,
		QuoteTimestamp: quote_timestamp,
		QuoteExtractedAt: time.Now(),
	}

	return curr	
}
