package company

import (
    "github.com/HoMuChen/go-twstock/domain"
)

type companyService struct {
    companySource   domain.CompanySource
}

func New(s domain.CompanySource) domain.CompanyService {
    return &companyService{
       companySource:  s,
    }
}

func (service *companyService) GetById(id string) (domain.Company, error) {
    return service.companySource.GetById(id)
}

func (service *companyService) List(from int, size int) ([]domain.Company, error) {
    return service.companySource.List(from, size)
}
