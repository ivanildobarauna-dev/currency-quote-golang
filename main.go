package main

import (
	"currency-quote/pkg/currencyquote"
)


func main() {
	client := currencyquote.ClientBuilder{
		CurrencyList: []string{"USD-BRL", "USD-EUR"},
	}

	client.GetHistoryCurrencyQuote(20230101)
	client.GetLastCurrencyQuote()
}