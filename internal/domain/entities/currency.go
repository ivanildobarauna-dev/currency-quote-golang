package entities

import (
	"errors"
	"fmt"
)

type Currency struct {
	Codes [2]string
}

func (c Currency) Init() (*Currency, error) {

	fmt.Println("Starting CurrencyCodes validation")
	for _, code := range c.Codes {
		if len(code) != 3 {
			return nil, errors.New("currency_code invalid" + code)
		}
	}

	fmt.Println("CurrencyCodes Validated! OK")
	return &c, nil

}
