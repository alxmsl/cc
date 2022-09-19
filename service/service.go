package service

import "github.com/alxmsl/cc/service/coinmarketcap"

type ServiceInterface interface {
	Convert(amount float64, from, to string) (float64, error)
}

func New() ServiceInterface {
	return coinmarketcap.New()
}
