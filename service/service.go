package service

import (
	"github.com/alxmsl/cc/converter"
	"github.com/alxmsl/cc/converter/coinmarketcap"
)

type ServiceInterface interface {
	Convert(amount float64, from, to string) (float64, error)
}

type Service struct {
	converter converter.Converter
}

func New() *Service {
	return &Service{
		converter: coinmarketcap.New(),
	}
}

func (s *Service) Convert(amount float64, from, to string) (float64, error) {
	return s.converter.Convert(amount, from, to)
}
