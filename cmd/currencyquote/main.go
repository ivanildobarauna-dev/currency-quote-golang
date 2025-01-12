package main

import (
	"currency-quote/pkg/currencyquote"
	"fmt"
)

func main() {
	currencyQuote, err := currencyquote.GetLastCurrencyQuote(
		"USD", "BRL",
	)

	if err == nil {
		fmt.Println(currencyQuote)
		return
	}
}
