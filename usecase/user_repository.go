package usecase

import (
    "github.com/maru0804/clearnACT.git/domain"
)

type UserRepository interface {
    Store(domain.User)
    Select() []domain.User
    Delete(id string)
}