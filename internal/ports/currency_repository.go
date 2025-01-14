package ports

import (
	"currency-quote/internal/domain/entities"
)

type ICurrencyRepository interface {
	GetLastCurrencyQuote(currencyList *entities.CurrencyPairList) (*[]entities.CurrencyQuote, error)
}
