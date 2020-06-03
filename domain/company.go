package domain

type Company struct {
    ID      string  `json:"id"`
    Code    string  `json:"code"`
    Name    string  `json:"name"`
}

type CompanyService interface {
    GetById(id string)              (Company, error)
    List(from int, size int)        ([]Company, error)
}

type CompanyRepository interface {
    GetById(id string)          (Company, error)
    List(from int, size int)    ([]Company, error)
    Add(company Company)        error
    Remove(company Company)     error
}

type CompanySource interface {
    GetById(id string)          (Company, error)
    List(from int, size int)    ([]Company, error)
}
