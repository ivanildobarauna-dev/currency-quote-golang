package adapters

import (
	"currency-quote/internal/domain/entities"
	"fmt"
)

type CurrencyRepository struct {
}

func (c CurrencyRepository) GetLastCurrencyQuote(currencyPair *entities.Currency) (*entities.CurrencyQuote, error) {
	fmt.Println("Starting CurrencyRepository / GetLastCurrencyQuote")

	currQuote := entities.NewCurrencyQuote(
		currencyPair,
		"undefined",
		currencyPair.Codes[0], // Base Currency
		currencyPair.Codes[1], // Quote Currency
		1231,
		5.433,
		3.44,
	)

	fmt.Println("Finished CurrencyRepository! OK")
	return currQuote, nil
}

func (c CurrencyRepository) GetHistoricalCurrencyQuote(currencyPair entities.Currency) (*entities.CurrencyQuote, error) {
	currQuote := entities.NewCurrencyQuote(
		&currencyPair,
		"undefined",
		currencyPair.Codes[0], // Base Currency
		currencyPair.Codes[1], // Quote Currency
		1231,
		5.433,
		3.44,
	)

	return currQuote, nil
}
