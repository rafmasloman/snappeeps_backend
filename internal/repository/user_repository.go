package repository

import (
	"github.com/rafmasloman/snappeeps_backend/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
}
