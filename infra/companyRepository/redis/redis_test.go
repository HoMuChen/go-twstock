package companyRepository_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-park/domain"
    redisRepo "github.com/HoMuChen/go-park/infra/companyRepository/redis"
)

func TestGetByNotExistID(t *testing.T) {
    repo := redisRepo.New()

    _, err := repo.GetById("nonExistKey")

    assert.Equal(t, err, domain.ErrNotFound, "err should be domain.ErrNotFound")
}

func TestList(t *testing.T) {
    repo := redisRepo.New()

    repo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    repo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    repo.Add(domain.Company{"2330", "tse_2330.tw", "TSMC"})

    companies, err := repo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 1)

    repo.Add(domain.Company{"2454", "tse_2454.tw", "發哥"})
    companies, err = repo.List(0, 2)
    assert.NoError(t, err)
    assert.Len(t, companies, 2)

    repo.Add(domain.Company{"2884", "tse_2884.tw", "玉山妹子"})
    companies, err = repo.List(0, 3)
    assert.NoError(t, err)
    assert.Len(t, companies, 3)
}

func TestAddAndGetById(t *testing.T) {
    repo := redisRepo.New()

    company := domain.Company{"2330", "tse_2330.tw", "TSMC"}

    err := repo.Add(company)
    assert.NoError(t, err)

    doc, err := repo.GetById("2330")
    assert.Equal(t, doc, company, "should be same company")
}

func TestRemove(t *testing.T) {
    repo := redisRepo.New()

    company := domain.Company{"2330", "tse_2330.tw", "TSMC"}

    err := repo.Add(company)
    assert.Equal(t, err, nil, "err should be nil")

    doc, err := repo.GetById("2330")
    assert.Equal(t, doc, company, "should be same company")

    err = repo.Remove(domain.Company{"2330", "tse_2330.tw", "TSMC"})
    assert.NoError(t, err)

    _, err = repo.GetById("2330")
    assert.Equal(t, err, domain.ErrNotFound, "err should be domain.ErrNotFound")
}
