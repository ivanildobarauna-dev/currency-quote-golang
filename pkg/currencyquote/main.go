package currencyquote

import (
	"currency-quote/internal/domain/entities"
	"fmt"
)

type ClientBuilder struct {
	Currency entities.Currency
}

func (c ClientBuilder) GetLastCurrencyQuote() {
	fmt.Println("Getting Last Quote")
	fmt.Println("CurrencyList: ", c.Currency.CurrencyList)
}

func (c ClientBuilder) GetHistoryCurrencyQuote(reference_date int) {
	fmt.Println("Getting History Quote")
	fmt.Println("CurrencyList: ", c.Currency.CurrencyList, "date: ", reference_date)
}