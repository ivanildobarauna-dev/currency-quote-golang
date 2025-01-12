package currencyquote

import (
	"currency-quote/internal/domain/entities"
)

func GetLastCurrencyQuote(baseCurrency string, quoteCurrency string) (*entities.CurrencyQuote, error) {
	currency, err := entities.NewCurrency([2]string{baseCurrency, quoteCurrency})

	if err != nil {
		return nil, err
	}

	currQuote := entities.NewCurrencyQuote(
		currency,
		"undefined",
		baseCurrency,
		quoteCurrency,
		1231,
		5.433,
		3.44,
	)

	return currQuote, nil
}
