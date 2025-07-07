package mapper

import (
	"github.com/rafmasloman/snappeeps_backend/internal/domain"
	"github.com/rafmasloman/snappeeps_backend/internal/entity"
)

func MapperUserResponse(user *domain.User) *entity.User {
	return &entity.User{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
