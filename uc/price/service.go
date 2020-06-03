package price

import (
    "sync"

    "github.com/HoMuChen/go-twstock/domain"
)

type priceService struct {
    priceSource       domain.PriceHttpSource
    priceRepo         domain.PriceRepository
    companyRepo       domain.CompanyRepository
}

func New(source domain.PriceHttpSource, pr domain.PriceRepository, cr domain.CompanyRepository) domain.PriceService {
    return &priceService{
       priceSource:  source,
       priceRepo:    pr,
       companyRepo:  cr,
    }
}

func (service *priceService) FetchRealtime(company domain.Company) (price domain.Price, err error) {
    return service.priceSource.FetchRealtime(company)
}

func (service *priceService) FetchRealtimeAll(from int, size int) (prices []domain.Price, err error) {
    companies, err := service.companyRepo.List(from, size)
    if err != nil {
        return
    }

    var wg sync.WaitGroup
    for _, company := range companies {
        wg.Add(1)

        go func(company domain.Company) {
            price, err := service.priceSource.FetchRealtime(company)
            if err != nil {
                wg.Done()
                return
            }

            prices = append(prices, price)
            wg.Done()
        }(company)
    }
    wg.Wait()

    return prices, nil
}
