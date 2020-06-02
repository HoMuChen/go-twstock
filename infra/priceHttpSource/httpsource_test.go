package httpsource_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-park/domain"
    httpsource "github.com/HoMuChen/go-park/infra/priceHttpSource"
)

func TestFetchRealtime(t *testing.T) {
    priceService := httpsource.New()

    price, err := priceService.FetchRealtime(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    assert.NoError(t, err)
    assert.NotEmpty(t, price)

    t.Log(price)
}

func TestFetchRealtimeBadInput(t *testing.T) {
    priceService := httpsource.New()

    price, err := priceService.FetchRealtime(domain.Company{"2330", "2330", "TSMC"})
    assert.Error(t, err)
    assert.Equal(t, price, domain.Price{})
}
