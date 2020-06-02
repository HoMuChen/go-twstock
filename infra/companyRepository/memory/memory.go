package companyRepository

import (
    "github.com/HoMuChen/go-park/domain"
)

type comanyRepository struct {
    companies map[string]domain.Company
}

func New() domain.CompanyRepository {
    return &comanyRepository{
        make(map[string]domain.Company),
    }
}

func (repo *comanyRepository) GetById(id string) (domain.Company, error) {
    company, found := repo.companies[id]

    if found {
        return company, nil
    }

    return company, domain.ErrNotFound
}

func (repo *comanyRepository) List(from int, size int) ([]domain.Company, error) {
    companies := make([]domain.Company, 0)

    count := 0
    for _, v := range repo.companies {
        if count < from {
            count++
            continue
        }
        companies = append(companies, v)

        if len(companies) > size {
            break
        }
    }

    return companies, nil
}

func (repo *comanyRepository) Add(company domain.Company) error {
    repo.companies[company.ID] = company

    return nil
}

func (repo *comanyRepository) Remove(company domain.Company) error {
    delete(repo.companies, company.ID)

    return nil
}
