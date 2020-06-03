package follow_test

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/HoMuChen/go-park/domain"
    repo    "github.com/HoMuChen/go-park/infra/companyRepository/memory"
    service "github.com/HoMuChen/go-park/uc/follow"
)

func TestFollow(t *testing.T) {
    memoryRepo := repo.New()
    followService := service.New(memoryRepo)

    err := followService.Follow(domain.Company{"1", "tse_1.tw", "台1"})
    assert.NoError(t, err)
}

func TestList(t *testing.T) {
    memoryRepo := repo.New()
    followService := service.New(memoryRepo)

    followService.Follow(domain.Company{"1", "tse_1.tw", "台1"})
    followService.Follow(domain.Company{"2", "tse_2.tw", "台2"})
    followService.Follow(domain.Company{"3", "tse_3.tw", "台3"})

    companies, err := followService.List(0, 5)
    assert.NoError(t, err)
    assert.Len(t, companies, 3)

    t.Log(companies)
}

