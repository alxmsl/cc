package coinmarketcap

import (
	"fmt"
	"os"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

type Service struct {
	client *cmc.Client
}

func New() *Service {
	return &Service{
		client: cmc.NewClient(&cmc.Config{
			ProAPIKey: os.Getenv("CMC_API_KEY"),
		}),
	}
}

func (s *Service) Convert(amount float64, from, to string) (r float64, err error) {
	res, err := s.client.Tools.PriceConversion(&cmc.ConvertOptions{
		Amount:  amount,
		Symbol:  from,
		Convert: to,
	})
	if err != nil {
		return
	}
	if _, ok := res.Quote[to]; !ok {
		return r, fmt.Errorf("coinmarketcap has not found the ticker '%s'", to)
	}
	r = res.Quote[to].Price
	return
}
