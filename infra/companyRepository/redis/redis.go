package companyRepository

import (
    "encoding/json"
    "github.com/go-redis/redis"

    "github.com/HoMuChen/go-park/domain"
)

type comanyRepository struct {
    namespace   string
    client      *redis.Client
}

func New() domain.CompanyRepository {
    return &comanyRepository{
        namespace:   "company:",
        client:      redis.NewClient(&redis.Options{
                         Addr:   "localhost:6379",
                     }),
    }
}

func (repo *comanyRepository) GetById(id string) (domain.Company, error) {
    val, err := repo.client.Get(repo.genKey(id)).Bytes()
    if err != nil {
        return domain.Company{}, domain.ErrNotFound
    }

    var company domain.Company
    err = json.Unmarshal(val, &company)
    if err != nil {
        return company, err
    }

    return company, nil
}

func (repo *comanyRepository) List(from int, size int) (companies []domain.Company, err error) {
    keys, err := repo.client.Keys(repo.genKey("*")).Result()

    var company domain.Company
    for _, key := range keys[from:] {
        val, _ := repo.client.Get(key).Bytes()

        json.Unmarshal(val, &company)
        companies = append(companies, company)

        if len(companies) == size {
            return companies, err
        }
    }

    return companies, err
}

func (repo *comanyRepository) Add(company domain.Company) error {
    encoded, err := json.Marshal(company)
    if err != nil {
        return err
    }

    err = repo.client.Set(repo.genKey(company.ID), encoded, 0).Err()
    if err != nil {
        return err
    }

    return nil
}

func (repo *comanyRepository) Remove(company domain.Company) (err error) {
    err = repo.client.Del(repo.genKey(company.ID)).Err()

    return
}

func (repo *comanyRepository) genKey(id string) string {
    return repo.namespace + id
}
