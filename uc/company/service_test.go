package company_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-park/domain"
    repo    "github.com/HoMuChen/go-park/infra/companyRepository/memory"
    source  "github.com/HoMuChen/go-park/infra/companySource/file"
    service "github.com/HoMuChen/go-park/uc/company"
)

func TestGetByID(t *testing.T) {
    memoryRepo := repo.New()
    fileSource := source.New("../../data/companies.csv")
    companyService := service.New(memoryRepo, fileSource)

    com, err := companyService.GetById("2330")
    assert.NoError(t, err)
    assert.Equal(t, com, domain.Company{"2330", "tse_2330.tw", "台積電"})
}

func TestList(t *testing.T) {
    memoryRepo := repo.New()
    fileSource := source.New("../../data/companies.csv")
    companyService := service.New(memoryRepo, fileSource)

    companies, err := companyService.List(0, 10)
    assert.NoError(t, err)
    assert.Len(t, companies, 10)
}

func TestFollow(t *testing.T) {
    memoryRepo := repo.New()
    fileSource := source.New("../../data/companies.csv")
    companyService := service.New(memoryRepo, fileSource)

    company, err := companyService.GetById("2330")

    err = companyService.Follow(company)
    assert.NoError(t, err)
}

func TestListFollow(t *testing.T) {
    memoryRepo := repo.New()
    fileSource := source.New("../../data/companies.csv")
    companyService := service.New(memoryRepo, fileSource)

    companyService.Follow(domain.Company{"1", "tse_1.tw", "台1"})
    companyService.Follow(domain.Company{"2", "tse_2.tw", "台2"})
    companyService.Follow(domain.Company{"3", "tse_3.tw", "台3"})

    companies, err := companyService.ListFollow(0, 5)
    assert.NoError(t, err)
    assert.Len(t, companies, 3)

    t.Log(companies)
}

