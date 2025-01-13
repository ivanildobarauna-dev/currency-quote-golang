package main

import (
	"currency-quote/pkg/currencyquote"
	"fmt"
)

func main() {
	fmt.Println("Calling cmd/currencyquote/main.go")
	currencyQuote, err := currencyquote.GetLastCurrencyQuote(
		"USD", "BRL",
	)

	if err == nil {
		fmt.Println(currencyQuote)
		fmt.Println("Finished cmd/currencyquote/main.go! OK")
		return
	}
}
