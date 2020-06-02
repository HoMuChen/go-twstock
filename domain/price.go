package domain

import "time"

type Price struct {
    Company      Company
    Time         time.Time
    Value        float64
}

type PriceService interface {
    FetchRealtime(company Company) (Price, error)
    FetchRealtimeAll(from int, to int) ([]Price, error)
}

type PriceRepository interface {
}

type PriceHttpSource interface {
    FetchRealtime(company Company) (Price, error)
}
