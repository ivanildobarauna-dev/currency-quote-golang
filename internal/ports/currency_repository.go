package ports

import (
	"currency-quote/internal/domain/entities"
)

type ICurrencyRepository interface {
	GetLastCurrencyQuote(currency *entities.Currency) (*entities.CurrencyQuote, error)
	//GetHistoricalCurrencyQuote(baseCurrency string, quoteCurrency string, referenceDate int) (*entities.CurrencyQuote, error)
}
