package domain

type FollowService interface {
    Follow(company Company)     error
    Unfollow(company Company)   error
    List(from int, size int)    ([]Company, error)
}
