package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafmasloman/snappeeps_backend/internal/dto"
	"github.com/rafmasloman/snappeeps_backend/internal/usecase"
)

type AuthHttpImpl struct {
	service usecase.AuthUsecase
}

func NewAuthHttp(service usecase.AuthUsecase) *AuthHttpImpl {
	return &AuthHttpImpl{service}
}

func (ctl AuthHttpImpl) Register(c *gin.Context) {
	var req dto.DTORegisterAuth

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	result, err := ctl.service.Register(&req)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  false,
			"message": "registrasi failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "registrasi berhasil",
		"result":  &result,
	})
}
