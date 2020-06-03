package companyRepository_test

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-twstock/domain"
    file "github.com/HoMuChen/go-twstock/infra/companyRepository/file"
)

const (
    filePath = "./test.csv"
)

func TestGetByNotExistID(t *testing.T) {
    fileRepo := file.New(filePath)

    _, err := fileRepo.GetById("2330")

    assert.Equal(t, err, domain.ErrNotFound, "err should be domain.ErrNotFound")

    os.Remove(filePath)
}

func TestList(t *testing.T) {
    fileRepo := file.New(filePath)

    fileRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    fileRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    fileRepo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})

    companies, err := fileRepo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 1)

    fileRepo.Add(domain.Company{"2454", "tse_2454.tw", "發哥"})
    companies, err = fileRepo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 2)

    fileRepo.Add(domain.Company{"2884", "tse_2884.tw", "玉山妹子"})
    companies, err = fileRepo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 3)

    os.Remove(filePath)
}

func TestAddAndGetById(t *testing.T) {
    fileRepo := file.New(filePath)

    company := domain.Company{"2330", "tse_2330.tw", "TSMC"}

    err := fileRepo.Add(company)
    assert.Equal(t, err, nil, "err should be nil")

    doc, err := fileRepo.GetById("2330")
    assert.Equal(t, doc, company, "should be same company")

    os.Remove(filePath)
}

func TestRemove(t *testing.T) {
    fileRepo := file.New(filePath)

    company := domain.Company{"2330", "tse_2330.tw", "TSMC"}

    err := fileRepo.Add(company)
    assert.Equal(t, err, nil, "err should be nil")

    doc, err := fileRepo.GetById("2330")
    assert.Equal(t, doc, company, "should be same company")

    err = fileRepo.Remove(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    assert.NoError(t, err)

    _, err = fileRepo.GetById("2330")
    assert.Equal(t, err, domain.ErrNotFound, "err should be domain.ErrNotFound")

    os.Remove(filePath)
}
