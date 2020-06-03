package companyRepository

import (
    "os"
    "bufio"
    "log"
    "strings"
    "fmt"

    "github.com/HoMuChen/go-twstock/domain"
)

type comanyRepository struct {
    companies map[string]domain.Company
    file      *os.File
}

func New(dataPath string) domain.CompanyRepository {
    companies := make(map[string]domain.Company)

    file, err := os.OpenFile(dataPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0655)
    if err != nil {
        log.Fatal(err)
    }

    reader := bufio.NewReader(file)

    for {
        line, _, err := reader.ReadLine()

        if err != nil {
            break
        }

        row := strings.Split(string(line), ",")

        if row[3] == "1" {
            companies[row[0]] = domain.Company{row[0], row[1], row[2]}
        }
        if row[3] == "-1" {
            delete(companies, row[0])
        }
    }

    return &comanyRepository{
        companies: companies,
        file:      file,
    }
}

func (repo *comanyRepository) GetById(id string) (domain.Company, error) {
    company, found := repo.companies[id]

    if found {
        return company, nil
    }

    return company, domain.ErrNotFound
}

func (repo *comanyRepository) List(from int, size int) (companies []domain.Company, err error) {
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

    return
}

func (repo *comanyRepository) Add(company domain.Company) (err error) {
    repo.companies[company.ID] = company

    _, err = repo.file.WriteString(fmt.Sprintf("%s,%s,%s,1\n", company.ID, company.Code, company.Name))

    return
}

func (repo *comanyRepository) Remove(company domain.Company) (err error) {
    delete(repo.companies, company.ID)

    _, err = repo.file.WriteString(fmt.Sprintf("%s,%s,%s,-1\n", company.ID, company.Code, company.Name))

    return
}
