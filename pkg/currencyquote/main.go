package currencyquote

import (
	"currency-quote/internal/adapters"
	"currency-quote/internal/domain/entities"
	"currency-quote/internal/domain/services"
	"fmt"
)

func GetLastCurrencyQuote(baseCurrency string, quoteCurrency string) (*entities.CurrencyQuote, error) {
	fmt.Println("Use Case Called")
	currency, err := entities.Currency{
		Codes: [2]string{baseCurrency, quoteCurrency},
	}.Init()

	if err != nil {
		return nil, err
	}

	currQuote, err := services.GetLastQuoteService(currency, adapters.CurrencyRepository{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Currency Quote Extracted", currQuote)

	return currQuote, nil
}
