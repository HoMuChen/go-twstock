package companyRepository_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-twstock/domain"
    memory "github.com/HoMuChen/go-twstock/infra/companyRepository/memory"
)

func TestGetByNotExistID(t *testing.T) {
    memoryRepo := memory.New()

    _, err := memoryRepo.GetById("2330")

    assert.Equal(t, err, domain.ErrNotFound, "err should be domain.ErrNotFound")
}

func TestList(t *testing.T) {
    memoryRepo := memory.New()

    memoryRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    memoryRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    memoryRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})

    companies, err := memoryRepo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 1)

    memoryRepo.Add(domain.Company{"2454", "tse_2454.tw", "發哥"})
    companies, err = memoryRepo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 2)

    memoryRepo.Add(domain.Company{"2884", "tse_2884.tw", "玉山妹子"})
    companies, err = memoryRepo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 3)
}

func TestAddAndGetById(t *testing.T) {
    memoryRepo := memory.New()

    company := domain.Company{"2330", "tse_2330.tw", "TSMC"}

    err := memoryRepo.Add(company)
    assert.Equal(t, err, nil, "err should be nil")

    doc, err := memoryRepo.GetById("2330")
    assert.Equal(t, doc, company, "should be same company")
}

func TestRemove(t *testing.T) {
    memoryRepo := memory.New()

    company := domain.Company{"2330", "tse_2330.tw", "TSMC"}

    err := memoryRepo.Add(company)
    assert.Equal(t, err, nil, "err should be nil")

    doc, err := memoryRepo.GetById("2330")
    assert.Equal(t, doc, company, "should be same company")

    err = memoryRepo.Remove(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    assert.NoError(t, err)

    _, err = memoryRepo.GetById("2330")
    assert.Equal(t, err, domain.ErrNotFound, "err should be domain.ErrNotFound")
}
