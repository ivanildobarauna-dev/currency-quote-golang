package services

import (
	"currency-quote/internal/domain/entities"
	"currency-quote/internal/ports"
	"fmt"
)

func GetLastQuoteService(c *entities.Currency, repo ports.ICurrencyRepository) (*entities.CurrencyQuote, error) {
	fmt.Println("Starting GetLastQuoteService")
	repositoryResponse, err := repo.GetLastCurrencyQuote(c)

	fmt.Println("GetLastQuoteService finished! OK")
	return repositoryResponse, err
}
