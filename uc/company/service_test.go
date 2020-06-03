package company_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-twstock/domain"
    source  "github.com/HoMuChen/go-twstock/infra/companySource/file"
    service "github.com/HoMuChen/go-twstock/uc/company"
)

func TestGetByID(t *testing.T) {
    fileSource := source.New("../../data/companies.csv")
    companyService := service.New(fileSource)

    com, err := companyService.GetById("2330")
    assert.NoError(t, err)
    assert.Equal(t, com, domain.Company{"2330", "tse_2330.tw", "台積電"})
}

func TestList(t *testing.T) {
    fileSource := source.New("../../data/companies.csv")
    companyService := service.New(fileSource)

    companies, err := companyService.List(0, 10)
    assert.NoError(t, err)
    assert.Len(t, companies, 10)
}

