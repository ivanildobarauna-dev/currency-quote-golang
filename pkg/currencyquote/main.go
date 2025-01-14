package currencyquote

import (
	"currency-quote/internal/adapters"
	"currency-quote/internal/domain/entities"
	"currency-quote/internal/domain/services"
	"fmt"
)

func GetLastCurrencyQuote(CurrencyPairCodes []string) (*[]entities.CurrencyQuote, error) {

	currencyPairList := entities.CurrencyPairList{
		Pairs: CurrencyPairCodes,
	}

	currQuote, err := services.GetLastQuoteService(&currencyPairList, adapters.CurrencyRepository{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Currency Quote Extracted", currQuote)

	return currQuote, nil
}
