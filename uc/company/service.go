package company

import (
    "github.com/HoMuChen/go-park/domain"
)

type companyService struct {
    companyRepo     domain.CompanyRepository
    companySource   domain.CompanySource
}

func New(c domain.CompanyRepository, s domain.CompanySource) domain.CompanyService {
    return &companyService{
       companyRepo:    c,
       companySource:  s,
    }
}

func (service *companyService) GetById(id string) (domain.Company, error) {
    return service.companySource.GetById(id)
}

func (service *companyService) List(from int, size int) ([]domain.Company, error) {
    return service.companySource.List(from, size)
}

func (service *companyService) Follow(company domain.Company) error {
    return service.companyRepo.Add(company)
}

func (service *companyService) ListFollow(from int, size int) ([]domain.Company, error) {
    return service.companyRepo.List(from, size)
}
