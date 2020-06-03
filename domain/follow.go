package domain

type FollowService interface {
    Follow(company Company)     error
    List(from int, size int)    ([]Company, error)
}
