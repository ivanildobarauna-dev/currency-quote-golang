package currencyquote

import "fmt"

type ClientBuilder struct {
	CurrencyList []string
}

func (c ClientBuilder) GetLastCurrencyQuote() {
	fmt.Println("Getting Last Quote")
	fmt.Println("CurrencyList: ", c.CurrencyList)
}

func (c ClientBuilder) GetHistoryCurrencyQuote(reference_date int) {
	fmt.Println("Getting History Quote")
	fmt.Println("CurrencyList: ", c.CurrencyList, "date: ", reference_date)
}