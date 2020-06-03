package memory

import (
    "github.com/HoMuChen/go-twstock/domain"
)

type priceRepository struct {
    prices map[string][]domain.Price
}

func New() domain.PriceRepository {
    return &priceRepository{
        make(map[string][]domain.Price),
    }
}

func (repo *priceRepository) Range() error {
    return nil
}
