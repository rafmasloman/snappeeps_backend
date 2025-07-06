package repository

import (
	"github.com/rafmasloman/snappeeps_backend/internal/entity"
)

type UserRepository interface {
	Create(*entity.User) error
	FindByUsername(username string)
}
