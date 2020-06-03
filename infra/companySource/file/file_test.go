package companySource_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-twstock/domain"
    _companySource "github.com/HoMuChen/go-twstock/infra/companySource/file"
)

func TestGetByID(t *testing.T) {
    companySource := _companySource.New("../../../data/companies.csv")

    com, err := companySource.GetById("2330")

    assert.NoError(t, err)
    assert.Equal(t, com, domain.Company{"2330", "tse_2330.tw", "台積電"} ,"err should be domain.ErrNotFound")
}

func TestGetByNonExistID(t *testing.T) {
    companySource := _companySource.New("../../../data/companies.csv")

    _, err := companySource.GetById("2330000")

    assert.Equal(t, err, domain.ErrNotFound)
}

func TestList(t *testing.T) {
    companySource := _companySource.New("../../../data/companies.csv")

    companies, err := companySource.List(0, 10)

    assert.NoError(t, err)
    assert.Len(t, companies, 10)

    t.Log(companies)
}
