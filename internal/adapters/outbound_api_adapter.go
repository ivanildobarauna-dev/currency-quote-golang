package adapters

import (
	"currency-quote/internal/domain/entities"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	URL                         = "https://economia.awesomeapi.com.br"
	ENDPOINT_AVALIABLE_PARITIES = "/json/available"
	ENDPOINT_LAST_COTATION      = "/last/"
	ENDPOINT_HISTORY_COTATION   = "/json/daily/"
	RETRY_TIME_SECONDS          = 2
	RETRY_ATTEMPTS              = 3
)

type CurrencyRepository struct {
}

func (c CurrencyRepository) GetLastCurrencyQuote(currencyList *entities.CurrencyPairList) (*[]entities.CurrencyQuote, error) {

	currencyList.Params = strings.Join(currencyList.Pairs, ",")

	resp, err := http.Get(URL + ENDPOINT_LAST_COTATION + currencyList.Params)

	if err != nil {
		fmt.Println("error requesting resource", err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error reading response body", err.Error())
	}

	fmt.Println(string(body))

	currQuotes := entities.NewCurrencyQuote(
		currencyList,
		"undefined",
		"USD",
		"BRL",
		1231,
		5.433,
		3.44,
	)
	return currQuotes, nil
}
