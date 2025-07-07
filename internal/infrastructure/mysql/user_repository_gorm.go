package mysql

import (
	"fmt"

	"github.com/rafmasloman/snappeeps_backend/internal/domain"
	"github.com/rafmasloman/snappeeps_backend/internal/entity"
	"github.com/rafmasloman/snappeeps_backend/internal/infrastructure/mapper"
	"github.com/rafmasloman/snappeeps_backend/internal/repository"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user *entity.User) (*entity.User, error) {
	// var userModel domain.User

	userModel := domain.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
	}

	if err := r.db.Create(&userModel).Error; err != nil {
		return nil, fmt.Errorf(`failed to create user %w`, err)
	}

	userCreateted := mapper.MapperUserResponse(&userModel)

	return userCreateted, nil
}

func (r *userRepositoryImpl) FindByUsername(email string) (*entity.User, error) {
	var userModel domain.User
	user := r.db.Where(`email = ?`, email).Find(&userModel)

	if user.Error != nil {
		return nil, fmt.Errorf(`user not found %w`, user.Error)
	}

	currentUser := mapper.MapperUserResponse(&userModel)

	return currentUser, nil
}
