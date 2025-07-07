package usecase

import (
	"fmt"

	"github.com/rafmasloman/snappeeps_backend/internal/dto"
	"github.com/rafmasloman/snappeeps_backend/internal/entity"
	"github.com/rafmasloman/snappeeps_backend/internal/repository"
	"github.com/rafmasloman/snappeeps_backend/internal/usecase/helpers"
	"github.com/rafmasloman/snappeeps_backend/utils"
)

type AuthUsecase interface {
	Register(data *dto.DTORegisterAuth) (*entity.User, error)
	Login(data *dto.DTOLoginAuth) (*string, error)
}

type AuthUsecaseImpl struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return AuthUsecaseImpl{repo: repo}
}

func (s AuthUsecaseImpl) Register(data *dto.DTORegisterAuth) (*entity.User, error) {

	hashPassword, err := helpers.HashPassword(data.Password)

	if err != nil {
		return nil, fmt.Errorf(`failed to registered password %w`, err)
	}

	registerData := entity.User{
		Firstname: data.Firstname,
		Lastname:  data.Lastname,
		Email:     data.Email,
		Password:  hashPassword,
	}

	result, _ := s.repo.Create(&registerData)

	return result, nil

}

func (s AuthUsecaseImpl) Login(data *dto.DTOLoginAuth) (*string, error) {

	findUser, err := s.repo.FindByUsername(data.Email)

	if err != nil {
		return nil, fmt.Errorf(`wrong type email or password %w`, err)
	}

	checkPassword := helpers.ComparePassword(findUser.Password, data.Password)

	if !checkPassword {
		return nil, fmt.Errorf(`wrong type email or password %w`, err)
	}

	token, _ := utils.GenerateJWT(findUser.ID)

	return &token, nil
}
