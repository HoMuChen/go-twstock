package companySource

import (
    "os"
    "io"
    "log"
    "encoding/csv"

    "github.com/HoMuChen/go-park/domain"
)

type comanySource struct {
    companies   map[string]domain.Company
    keys        []string
}

func New(dataPath string) domain.CompanySource {
    companies := make(map[string]domain.Company)
    keys := make([]string, 0)

    file, err := os.Open(dataPath)
    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(file)
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        keys = append(keys, record[0])
        companies[record[0]] = domain.Company{record[0], record[1], record[2]}
    }

    return &comanySource{
        companies,
        keys,
    }
}

func (repo *comanySource) GetById(id string) (domain.Company, error) {
    company, found := repo.companies[id]

    if found {
        return company, nil
    }

    return company, domain.ErrNotFound
}

func (repo *comanySource) List(from int, size int) ([]domain.Company, error) {
    companies := make([]domain.Company, 0)

    count := 0
    for _, k := range repo.keys {
        if (count < from) {
            count++
            continue
        }
        companies = append(companies, repo.companies[k])

        if len(companies) >= size {
            break
        }
    }

    return companies, nil
}
