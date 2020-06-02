package price_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-park/domain"
    comRepo    "github.com/HoMuChen/go-park/infra/companyRepository/memory"
    memory     "github.com/HoMuChen/go-park/infra/priceRepository/memory"
    httpsource "github.com/HoMuChen/go-park/infra/priceHttpSource"
    service    "github.com/HoMuChen/go-park/uc/price"
)

func TestFetchRealtime(t *testing.T) {
    priceRepo := memory.New()
    companyRepo := comRepo.New()
    priceHttpSource := httpsource.New()
    priceService := service.New(priceHttpSource, priceRepo, companyRepo)


    price, err := priceService.FetchRealtime(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    assert.NoError(t, err)
    assert.NotEmpty(t, price)
}

func TestFetchRealtimeBadInput(t *testing.T) {
    priceRepo := memory.New()
    companyRepo := comRepo.New()
    priceHttpSource := httpsource.New()
    priceService := service.New(priceHttpSource, priceRepo, companyRepo)

    price, err := priceService.FetchRealtime(domain.Company{"2330", "2330", "TSMC"})
    assert.Error(t, err)
    assert.Equal(t, price, domain.Price{})
}

func TestFetchRealtimeAll(t *testing.T) {
    priceRepo := memory.New()
    companyRepo := comRepo.New()
    priceHttpSource := httpsource.New()
    priceService := service.New(priceHttpSource, priceRepo, companyRepo)

    companyRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    companyRepo.Add(domain.Company{"2884", "tse_2884.tw", "ESun"})

    prices, err := priceService.FetchRealtimeAll(0, 2)
    assert.NoError(t, err)
    assert.Len(t, prices, 2)

    t.Log(prices)
}
