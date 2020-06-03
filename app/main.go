package main

import (
    "fmt"

    _companyService "github.com/HoMuChen/go-twstock/uc/company"
    _companyRepo    "github.com/HoMuChen/go-twstock/infra/companyRepository/memory"
    _companySource  "github.com/HoMuChen/go-twstock/infra/companySource/file"

    _priceService   "github.com/HoMuChen/go-twstock/uc/price"
    _priceRepo      "github.com/HoMuChen/go-twstock/infra/priceRepository/memory"
    _priceSouce     "github.com/HoMuChen/go-twstock/infra/priceHttpSource"
)

func main() {
    comapnyRepo := _companyRepo.New()
    companySource := _companySource.New("./data/companies.csv")
    companyService := _companyService.New(comapnyRepo, companySource)

    priceRepo := _priceRepo.New()
    priceHttpSource := _priceSouce.New()
    priceService := _priceService.New(priceHttpSource, priceRepo, comapnyRepo)

    com, _ := companyService.GetById("2330")
    companyService.Follow(com)
    com, _ = companyService.GetById("2454")
    companyService.Follow(com)
    com, _ = companyService.GetById("3105")
    companyService.Follow(com)
    com, _ = companyService.GetById("2884")
    companyService.Follow(com)

    prices, err := priceService.FetchRealtimeAll(0, 5)
    if err != nil {
        fmt.Println(err)
    }

    for _, price := range prices {
        fmt.Println(price.Company.ID, price.Company.Name, price.Value)
    }
}
