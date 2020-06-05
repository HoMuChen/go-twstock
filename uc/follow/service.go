package follow

import (
    "github.com/HoMuChen/go-twstock/domain"
)

type followService struct {
    companyRepo     domain.CompanyRepository
}

func New(c domain.CompanyRepository) domain.FollowService {
    return &followService{
       companyRepo:    c,
    }
}

func (service *followService) Follow(company domain.Company) error {
    return service.companyRepo.Add(company)
}

func (service *followService) Unfollow(company domain.Company) error {
    return service.companyRepo.Remove(company)
}

func (service *followService) List(from int, size int) ([]domain.Company, error) {
    return service.companyRepo.List(from, size)
}
